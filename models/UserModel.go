package models

import (
	"time"

	"gorm.io/gorm"
)

// indicando gorm.Model le decimos a gorm que me convierta en tabla la siguiente structura

type User struct {
	gorm.Model

	FirstName       string `gorm:"size:255; not null"`
	LastName        string `gorm:"size:255; not null"`
	Email           string `gorm:"size:100; unique; not null"`
	Age             int    `gorm:"not null"`
	Products        []Product
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	Purchases       int
	Country         string `gorm:"size:120; default:'unknown'"`
	Role            string `gorm:"type:enum('user','seller','unknown'); default:'unknown'; not null"`
	Gender          string `gorm:"type:enum('male','female','unknown'); default:'unknown'; not null"`
	IsAuthenticated bool
	IsBanned        bool
}

type UserSerializer struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Products  []Product `json:"products"`
	Purchases int       `json:"purcharses"`
	Country   string    `json:"country"`
	Role      string    `json:"role"`
	Gender    string    `json:"gender"`
}
