package utils

import "github.com/PaulWaldo/gomoney/pkg/domain/models"

type PaginatedResponse struct {
	Data  []models.Transaction
	Count int64
	HasNext bool
}

