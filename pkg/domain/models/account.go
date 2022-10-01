package models

import (
	"errors"
)

type AccountType struct {
	Slug string
}

func (r AccountType) String() string {
	return r.Slug
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
	case Checking.Slug:
		return Checking, nil
	case Savings.Slug:
		return Savings, nil
	case CreditCard.Slug:
		return CreditCard, nil
	}

	return Unknown, errors.New("unknown account type: " + s)
}

type Account struct {
	ID              int64
	Name            string
	Type            string
	Memo            string
	Routing         string
	AccountNumber   string
	Hidden          bool
	NetWorthInclude bool
	BudgetInclude   bool
	Transactions    []Transaction
}
