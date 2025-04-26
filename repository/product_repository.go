package repository

import (
	"context"
	"github.com/ltvinh9899/soa_test/model"
	"github.com/ltvinh9899/soa_test/dto"
	"gorm.io/gorm"
	"time"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetByID(ctx context.Context, id uint) (model.Product, error) {
	var product model.Product
	err := r.db.WithContext(ctx).
		Preload("Categories").
		First(&product, id).Error
	return product, err
}

func (r *ProductRepository) GetPaginated(ctx context.Context, page, limit int) ([]model.Product, error) {
	var products []model.Product
	offset := (page - 1) * limit
	
	err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Preload("Categories").
		Find(&products).Error
		
	return products, err
}

func (r *ProductRepository) Create(ctx context.Context, product *model.Product) (uint, error) {
	err := r.db.WithContext(ctx).Create(&product).Error
	return product.ID, err
}

// repositories/product_repository.go
func (r *ProductRepository) FilterProducts(ctx context.Context, filter dto.ProductFilter, page, limit int) ([]model.Product, error) {
    query := r.db.WithContext(ctx).Model(&model.Product{})

    // Search by name
    if filter.SearchQuery != "" {
		if filter.Type == "name" {
        	query = query.Where("name ILIKE ?", "%"+filter.SearchQuery+"%")
    	}
		if filter.Type == "description" {
			query = query.Where("description ILIKE ?", "%"+filter.SearchQuery+"%")
		}
	}

    // Status filter
    if filter.Status != "" {
        query = query.Where("status = ?", filter.Status)
    }

	query = query.Where("deleted_at IS NULL")

    // Pagination
    offset := (page - 1) * limit
    var products []model.Product
    err := query.Offset(offset).Limit(limit).Preload("Categories").Find(&products).Error

    return products, err
}

func (r *ProductRepository) Update(ctx context.Context, productId uint, updates map[string]interface{}, categoryIds []uint) error {
	err := r.db.WithContext(ctx).
		Model(&model.Product{}).
		Where("id = ?", productId).
		Updates(updates).Error
	
	if err != nil {
		return err
	}

	if categoryIds != nil {
		if err := r.db.WithContext(ctx).
			Where("product_id = ?", productId).
			Delete(&model.ProductCategory{}).Error; err != nil {
			return err
		}
		for _, catId := range categoryIds {
			if err := r.db.WithContext(ctx).Create(&model.ProductCategory{
				ProductID:  productId,
				CategoryID: catId,
			}).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *ProductRepository) SoftDelete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&model.Product{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
}

func (r *ProductRepository) GetCategoryStats(ctx context.Context) ([]dto.CategoryStat, error) {
	var stats []dto.CategoryStat

	// Single query to get all category counts
	err := r.db.WithContext(ctx).
		Model(&model.Product{}).
		Select("categories.name as category_name, count(*) as product_count").
		Joins("JOIN product_categories ON products.id = product_categories.product_id").
		Joins("JOIN categories ON product_categories.category_id = categories.id").
		Where("products.deleted_at IS NULL").
		Group("categories.name").
		Scan(&stats).Error

	return stats, err
}