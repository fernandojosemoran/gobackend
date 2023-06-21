package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	ID          int       `gorm:"primaryKey; autoIncrement"`
	Name        string    `gorm:"size:255; not null"`
	Category    string    `gorm:"type:enum('snacks','condiments','sauces','seafood','pet supplies','canned foods','Grains','beverages','personal care products','unknown'); default:'unknown'"`
	Price       float64   `gorm:"not null"`
	Discount    bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Owner       User      `gorm:"foreignKey:ID"`
	Sales       int       `gorm:"type:int; default:0"`
	Sensitivity string    `gorm:"not null; default:'unknown'"`
}

type ProductZerializer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Discount    bool   `json:"discount"`
	Owner       string `json:"owner"`
	Sales       int    `json:"sales"`
	Sensitivity string `json:"sensitivity"`
}
