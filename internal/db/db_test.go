package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func Test_NewSqliteInMemoryServicesPopulatesDatabase(t *testing.T) {
	services, _, err := NewSqliteInMemoryServices(&gorm.Config{}, true)
	require.NoError(t, err)
	accounts, err := services.Account.List()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(accounts))
	assert.Equal(t, 2, len(accounts[0].Transactions))
}
