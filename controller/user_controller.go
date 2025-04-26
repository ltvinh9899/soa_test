package controller

import (
	"net/http"

	"github.com/ltvinh9899/soa_test/dto"
	"github.com/ltvinh9899/soa_test/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Register(ctx *gin.Context) {
	var request dto.RegisterInput

	// Bind và validate input
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gọi service đăng ký
	userID, err := c.userService.Register(ctx, request.Email, request.Password, request.FullName, request.Username, request.Role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Trả về response
	ctx.JSON(http.StatusCreated, gin.H{
		"error_flag": 0,
		"message":     "success",
		"id":       userID,
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	var request dto.LoginInput

	// Validate input
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gọi service đăng nhập
	token, _, err := c.userService.Login(ctx, request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Trả về token và thông tin user
	ctx.JSON(http.StatusOK, gin.H{
		"error_flag": 0,
		"message":     "success",
		"token":    token,
	})
}