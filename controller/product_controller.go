package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ltvinh9899/soa_test/dto"
	"github.com/ltvinh9899/soa_test/model"
	"github.com/ltvinh9899/soa_test/service"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) GetProducts(ctx *gin.Context) {
	filter := dto.ProductFilter{
        Status:      ctx.Query("status"),
		Type: 	  ctx.Query("type"),
        SearchQuery: ctx.Query("search"),
    }

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))

	products, err := c.service.GetProducts(ctx,filter, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error_flag": 1, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) GetProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	
	product, err := c.service.GetProduct(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error_flag": 1, "message": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error_flag": 0, "message": "success", "data": product})
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var request dto.CreateProductInput
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error_flag": 1, "message": "Invalid input"})
		return
	}

	product := model.Product{
		Name:          request.Name,
		Description:   request.Description,
		Price:         request.Price,
		StockQuantity: request.StockQuantity,
	}

	for _, catID := range request.CategoryIDs {
		product.Categories = append(product.Categories, model.Category{ID: catID})
	}

	productId, createErr := c.service.CreateProduct(ctx, &product)
	if createErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error_flag": 1, "message": createErr.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"error_flag": 1, "message": "Product created successfully", "product_id": productId})
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	
	var input struct {
		Name          *string  `json:"name"`
		Description   *string  `json:"description"`
		Price         *float64 `json:"price"`
		StockQuantity *int     `json:"stock_quantity"`
		CategoryIDs   []uint   `json:"category_ids"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error_flag": 1, "message": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if input.Name != nil {
		updates["name"] = *input.Name
	}
	// Thêm các trường khác tương tự...

	if err := c.service.UpdateProduct(ctx, uint(id), updates, input.CategoryIDs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error_flag": 1, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error_flag": 0, "message": "Product updated successfully"})
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	
	if err := c.service.DeleteProduct(ctx, uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error_flag": 1, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error_flag": 0, "message": "Product deleted"})
}

func (c *ProductController) GetDashboard(ctx *gin.Context) {
	data, err := c.service.GetDashboardData(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}