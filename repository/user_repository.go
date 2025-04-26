package repository

import (
	"context"
	"github.com/ltvinh9899/soa_test/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Tạo user mới
func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// Kiểm tra email tồn tại
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}

// Kiểm tra usenname tồn tại
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return user, err
}