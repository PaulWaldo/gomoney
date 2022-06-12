package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Type string
	Name string
}
