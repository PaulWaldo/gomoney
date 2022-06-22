package domain

import (
	"github.com/PaulWaldo/gomoney/internal/db/models"
)

// type AccountType struct {
// 	slug string
// }

// func (r AccountType) String() string {
// 	return r.slug
// }

// const (
// 	AccountTypeChecking   = "checking"
// 	AccountTypeSavings    = "savings"
// 	AccountTypeCreditCard = "creditCard"
// )

// var (
// 	Unknown    = AccountType{""}
// 	Checking   = AccountType{AccountTypeChecking}
// 	Savings    = AccountType{AccountTypeSavings}
// 	CreditCard = AccountType{AccountTypeCreditCard}
// )

// func AccountTypeFromString(s string) (AccountType, error) {
// 	switch s {
// 	case Checking.slug:
// 		return Checking, nil
// 	case Savings.slug:
// 		return Savings, nil
// 	case CreditCard.slug:
// 		return CreditCard, nil
// 	}

// 	return Unknown, errors.New("unknown account type: " + s)
// }

// type AccountIDType uint
// type Account struct {
// 	ID          AccountIDType `json:"id"`
// 	Name        string        `json:"payee"`
// 	AccountType AccountType   `json:"accountType"`
// }

// func NewAccount(name string, accountType AccountType) Account {
// 	return Account{Name: name, AccountType: accountType}
// }

type AccountSvc interface {
	Create(name string, accountType models.AccountType) (uint, error)
	Get(id uint) (*models.Account, error)
	List() ([]*models.Account, error)
}

// type AccountDB interface {
// 	Create(name string, accountType AccountType) (AccountIDType, error)
// 	Get(id AccountIDType) (*Account, error)
// 	List()([]*Account, error)
// }
