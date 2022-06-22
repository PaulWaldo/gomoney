package models

import (
	"errors"

	"gorm.io/gorm"
)

type AccountType struct {
	slug string
}

func (r AccountType) String() string {
	return r.slug
}

const (
	AccountTypeChecking   = "checking"
	AccountTypeSavings    = "savings"
	AccountTypeCreditCard = "creditCard"
)

var (
	Unknown    = AccountType{""}
	Checking   = AccountType{AccountTypeChecking}
	Savings    = AccountType{AccountTypeSavings}
	CreditCard = AccountType{AccountTypeCreditCard}
)

func AccountTypeFromString(s string) (AccountType, error) {
	switch s {
	case Checking.slug:
		return Checking, nil
	case Savings.slug:
		return Savings, nil
	case CreditCard.slug:
		return CreditCard, nil
	}

	return Unknown, errors.New("unknown account type: " + s)
}

// type AccountIDType uint
// type Account struct {
// 	ID          AccountIDType `json:"id"`
// 	Name        string        `json:"payee"`
// 	AccountType AccountType   `json:"accountType"`
// }

type Account struct {
	gorm.Model
	Name string
	Type AccountType
}
