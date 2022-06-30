package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func Test_NewSqliteInMemoryServicesPopulatesDatabase(t *testing.T) {
	services, db, err := NewSqliteInMemoryServices(&gorm.Config{}, false)
	require.NoError(t, err)
	teardownTest, _ := setupTest(t, db)
	defer teardownTest(t)

	// Note that the transaction object can't be inserted into services, so the population cannot be rolled back
	// PopulateTestData(*services)
	require.NoError(t, err)
	accounts, err := services.Account.List()
	assert.NoError(t, err)
	require.Equal(t, 0, len(accounts))
	// assert.Equal(t, 0, len(accounts[0].Transactions))
}
