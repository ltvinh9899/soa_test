package model

type Category struct {
  ID          uint      `gorm:"primaryKey" json:"id"`
  Name        string    `gorm:"not null"     json:"name"`
  Description string    `json:"description,omitempty"`
}