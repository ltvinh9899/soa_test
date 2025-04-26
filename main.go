package main

import (
	"github.com/ltvinh9899/soa_test/config"
	"github.com/ltvinh9899/soa_test/controller"

	"github.com/ltvinh9899/soa_test/middleware"
	"github.com/ltvinh9899/soa_test/model"
	"github.com/ltvinh9899/soa_test/repository"
	"github.com/ltvinh9899/soa_test/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db := config.InitDB(cfg)
	runMigrations(db)

	// Initialize repositories
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	productService := service.NewProductService(productRepo)
	userService := service.NewUserService(userRepo)
	

	// Initialize controllers
	productController := controller.NewProductController(productService)
	userController := controller.NewUserController(userService)

	// Setup router
	router := gin.Default()
	router.POST("api/user/register", userController.Register)
	router.POST("api/user/login", userController.Login)

	// Protected routes
	authRoutes := router.Group("/api").Use(middleware.JWTAuth(cfg.JWTSecret))

	authRoutes.GET("/products", productController.GetProducts)
	authRoutes.GET("/product/:id", productController.GetProduct)
    authRoutes.POST("/product", productController.CreateProduct)
    authRoutes.PUT("/product/:id", productController.UpdateProduct)
	authRoutes.DELETE("/product/:id", productController.DeleteProduct)
	authRoutes.GET("/dashboard", middleware.AdminAccess(), productController.GetDashboard)

	// Start server
	router.Run(":" + cfg.Port)
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Category{},
		&model.ProductCategory{},
	)
}