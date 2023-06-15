package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	ID          int    `gorm:"primaryKey; autoincrement"`
	name        string `gorm:"NOT NULL"`
	category    string `gorm:"NOT NULL"`
	price       string `gorm:"NOT NULL"`
	discount    bool   `gorm:"DEFAULT:false"`
	create_at   string `gorm:"TIMESTAMP; DEFAULT:CURRENT_TIMESTAMP"`
	update_at   string `gorm:"TIMESTAMP; DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	owner       int    `gorm:"UNIQUE; DEFAULT:UNKNOW"`
	sales       int    `gorm:"DEFAULT:0"`
	sensitivity string `gorm:"NOT NULL; DEFAULT:UNKNOW"`
}
