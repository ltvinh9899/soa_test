package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string  `gorm:"size:255;not null"`
	Description   string  `gorm:"type:text"`
	Price         float64 `gorm:"type:decimal(10,2)"`
	StockQuantity int     `gorm:"default:0"`
	Status        string  `gorm:"size:50;default:'active'"`
	Categories    []Category `gorm:"many2many:product_categories;"`
	// SearchVector  string `gorm:"type:tsvector"`
}

// func (p *Product) BeforeSave(tx *gorm.DB) error {
// 	// Update search vector for full-text search
// 	tx.Exec(`
// 		UPDATE products 
// 		SET search_vector = to_tsvector('english', name || ' ' || description)
// 		WHERE id = ?
// 	`, p.ID)
// 	return nil
// }