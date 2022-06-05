package models

import (
	"gorm.io/gorm"
)

type AccountType struct {
	gorm.Model
	Type string
}

type Account struct {
	gorm.Model
	Name        string
	AccountType AccountType
}
