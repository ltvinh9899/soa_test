package service

import (
	"context"
	"errors"
	"time"
	"github.com/ltvinh9899/soa_test/model"
	"github.com/ltvinh9899/soa_test/repository"
	"github.com/ltvinh9899/soa_test/config"
	"golang.org/x/crypto/bcrypt"
	jwt "github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Đăng ký user mới
func (s *UserService) Register(ctx context.Context, email, password, fullName, username, role string) (uint, error) {
	// Kiểm tra email đã tồn tại
	if _, err := s.userRepo.GetByEmail(ctx, email); err == nil {
		return 0, errors.New("email already exists")
	}

	// Kiểm tra username đã tồn tại
	if _, err := s.userRepo.GetByUsername(ctx, email); err == nil {
		return 0, errors.New("email already exists")
	}

	// Tạo user mới với password đã hash
	user, err := model.NewUser(email, password, fullName, username, role)
	if err != nil {
		return 0, err
	}

	// Lưu vào database
	if err := s.userRepo.Create(ctx, user); err != nil {
		return 0, err
	}

	return user.ID, nil // Trả về ID số (uint)
}

func (s *UserService) Login(ctx context.Context, username, password string) (string, *model.User, error) {
	// 1. Tìm user bằng email
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", nil, errors.New("user or password is incorrect")
	}

	// 2. Kiểm tra password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, errors.New("user or password is incorrect")
	}

	// 3. Tạo JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,       // User ID (uint)
		"role": user.Role,     // Phân quyền
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Hết hạn sau 24h
	})

	tokenString, err := token.SignedString([]byte(config.LoadConfig().JWTSecret))
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	return tokenString, &user, nil
}