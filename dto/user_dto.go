package dto

type LoginInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"full_name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
}