package pdc_model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string
	Price       int64
	Stock       int64
	Image       string
	Description string
	ActivityID  int64 `gorm:"foreignKey:ID"`
}
