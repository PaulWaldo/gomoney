package domain

import "github.com/PaulWaldo/gomoney/pkg/domain/models"

type AccountSvc interface {
	Create(name string, accountType string) (uint, error)
	Get(id uint) (models.Account, error)
	List() ([]models.Account, error)
}


type Services struct {
	Account AccountSvc
}

