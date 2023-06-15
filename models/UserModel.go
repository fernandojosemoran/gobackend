package models

import "gorm.io/gorm"

// indicando gorm.Model le decimos a gorm que me convierta en tabla la siguiente structura

type User struct {
	gorm.Model

	ID               int
	FirstName        string
	LastName         string
	Email            string
	age              int
	Products         []string
	Crete_at         string
	Update_at        string
	Purchases        int
	Country          string
	role             string
	Gender           string
	is_authenticated bool
	is_banned        bool
}
