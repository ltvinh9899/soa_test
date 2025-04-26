package model

import (
	"gorm.io/gorm"
)

type ProductCategory struct {
    ProductID  uint `gorm:"primaryKey"`
    CategoryID uint `gorm:"primaryKey"`
}

// Thiết lập quan hệ
func (ProductCategory) BeforeCreate(db *gorm.DB) error {
    return db.SetupJoinTable(&Product{}, "Categories", &ProductCategory{})
}