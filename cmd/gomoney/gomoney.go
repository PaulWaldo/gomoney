package main

import (
	"github.com/PaulWaldo/gomoney/internal/application"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

func main() {
	application.RunApp(&application.AppData{Accounts: []models.Account{}})
}
