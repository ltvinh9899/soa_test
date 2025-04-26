package service

import (
	"context"

	"github.com/ltvinh9899/soa_test/dto"
	"github.com/ltvinh9899/soa_test/model"
	"github.com/ltvinh9899/soa_test/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProduct(ctx context.Context, id uint) (model.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) GetProducts(ctx context.Context, filter dto.ProductFilter, page, limit int) ([]model.Product, error) {
	return s.repo.FilterProducts(ctx, filter, page, limit)
}

func (s *ProductService) CreateProduct(ctx context.Context, product *model.Product) (uint, error) {
	return s.repo.Create(ctx, product)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id uint, updates map[string]interface{}, categoryIds []uint) error {
	if err := s.repo.Update(ctx, id, updates, categoryIds); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uint) error {
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		return err
	}
	return s.repo.SoftDelete(ctx, id)
}

func (s *ProductService) GetDashboardData(ctx context.Context) (*dto.DashboardResponse, error) {
	stats, err := s.repo.GetCategoryStats(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.DashboardResponse{
		CategoryStats: stats,
	}, nil
}