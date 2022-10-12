package db

import (
	"testing"
	"time"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TransactionServiceTestSuite struct {
	suite.Suite
	Services *domain.Services
	TxSvc    domain.TransactionSvc
	Accounts []models.Account
}

func (suite *TransactionServiceTestSuite) SetupTest() {
	services, err := NewSqliteInMemoryServices(migrationDir, false)
	require.NoError(suite.T(), err)
	suite.TxSvc = services.Transaction
	suite.Services = services
	suite.Accounts = []models.Account{{Name: "acct1"}, {Name: "acct2"}}

	for i := range suite.Accounts {
		err = suite.Services.Account.Create(&suite.Accounts[i])
		require.NoError(suite.T(), err)
	}
}

func (suite *TransactionServiceTestSuite) TearDownTest() {
	suite.Services.Db.Close()
}

func TestTransactionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionServiceTestSuite))
}

func (suite *TransactionServiceTestSuite) Test_transactionSvc_Create() {
	want := &models.Transaction{
		Payee:     "payee",
		Type:      "t",
		Amount:    123.45,
		Memo:      "memo",
		Date:      time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		AccountID: suite.Accounts[0].ID,
	}
	err := suite.TxSvc.Create(want)
	require.NoError(suite.T(), err)

	got, err := suite.TxSvc.Get(want.AccountID)
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), want.Payee, got.Payee)
	assert.Equal(suite.T(), want.Type, want.Type)
	assert.Equal(suite.T(), want.Amount, want.Amount)
	assert.Equal(suite.T(), want.Memo, want.Memo)
	assert.Equal(suite.T(), want.Date, want.Date)
	assert.Equal(suite.T(), want.AccountID, want.AccountID)
}

func (suite *TransactionServiceTestSuite) Test_transactionSvc_Create_DetectsForeignConstraint() {
	want := &models.Transaction{
		Payee:     "payee",
		Type:      "t",
		Amount:    123.45,
		Memo:      "memo",
		Date:      time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		AccountID: 666,
	}
	err := suite.TxSvc.Create(want)
	require.Error(suite.T(), err)
	assert.Equal(suite.T(), "FOREIGN KEY constraint failed", err.Error())
}

func (suite *TransactionServiceTestSuite) Test_transactionSvc_List() {
	transactions := []models.Transaction{
		{
			Payee:     "p1",
			Type:      "t1",
			Amount:    1.1,
			Memo:      "m1",
			Date:      time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			AccountID: suite.Accounts[0].ID,
		},
		{
			Payee:     "p2",
			Type:      "t2",
			Amount:    1.1,
			Memo:      "m2",
			Date:      time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC),
			AccountID: suite.Accounts[0].ID,
		},
		{
			Payee:     "p2",
			Type:      "t2",
			Amount:    1.1,
			Memo:      "m2",
			Date:      time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC),
			AccountID: suite.Accounts[0].ID,
		},
	}

	saved := []models.Transaction{}
	for _, tx := range transactions {
		err := suite.TxSvc.Create(&tx)
		require.NoError(suite.T(), err)
		saved = append(saved, tx)
	}

	loaded, err := suite.TxSvc.List()
	require.NoError(suite.T(), err)
	require.EqualValues(suite.T(), len(transactions), len(loaded))
	for i := range loaded {
		want := saved[i]
		got := loaded[i]
		assert.Equal(suite.T(), want, got)
		assert.Equal(suite.T(), want.ID, got.ID)
		assert.Equal(suite.T(), want.Payee, got.Payee)
		assert.Equal(suite.T(), want.Amount, got.Amount)
		assert.Equal(suite.T(), want.Memo, got.Memo)
		assert.Equal(suite.T(), want.Date, got.Date)
		assert.Equal(suite.T(), want.AccountID, got.AccountID)
	}
}

func (suite *TransactionServiceTestSuite) Test_transactionSvc_ListByAccount_ReturnsOnlySelectedTransactions() {
	txns := []models.Transaction{
		// Account 1
		{Payee: "acct1payee1", AccountID: suite.Accounts[0].ID},
		{Payee: "acct1payee2", AccountID: suite.Accounts[0].ID},
		// Account 2
		{
			Payee:     "p1",
			Type:      "t1",
			Amount:    1.1,
			Memo:      "m1",
			Date:      time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			AccountID: suite.Accounts[1].ID,
		},
		{
			Payee:     "p2",
			Type:      "t2",
			Amount:    1.1,
			Memo:      "m2",
			Date:      time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC),
			AccountID: suite.Accounts[1].ID,
		},
	}
	for i := range txns {
		err := suite.TxSvc.Create(&txns[i])
		require.NoError(suite.T(), err)
	}

	got, err := suite.TxSvc.ListByAccount(suite.Accounts[1].ID)
	require.NoError(suite.T(), err)
	require.EqualValuesf(suite.T(), 2, len(got), "expecting 2 transactions for account ID accounts[1].ID, got %d", len(got))
	assert.Equal(suite.T(), txns[2].Payee, got[0].Payee, "Expecting retrieved payee to be %s, got %s", txns[2].Payee, got[0].Payee)
	assert.Equal(suite.T(), txns[3].Payee, got[1].Payee, "Expecting retrieved payee to be %s, got %s", txns[3].Payee, got[1].Payee)
	assert.Equal(suite.T(), 1.1, got[0].Balance)
	assert.Equal(suite.T(), 2.2, got[1].Balance)
}

func (suite *TransactionServiceTestSuite) Test_transactionSvc_Update() {
	saved := models.Transaction{
		Payee:     "p1",
		Type:      models.Checking.Slug,
		Memo:      "m1",
		AccountID: suite.Accounts[0].ID,
	}
	err := suite.TxSvc.Create(&saved)
	assert.NoError(suite.T(), err, "Unable to save initial record: %s")

	saved.Payee = "p2"
	saved.Memo = "m2"
	saved.Type = models.Savings.Slug
	err = suite.TxSvc.Update(&saved)
	require.NoError(suite.T(), err, "Update threw error: %s", err)

	loaded, err := suite.TxSvc.Get(saved.AccountID)
	require.NoError(suite.T(), err, "Result load threw error: %s", err)
	assert.Equal(suite.T(), loaded.Payee, "p2")
	assert.Equal(suite.T(), loaded.Memo, "m2")
	assert.Equal(suite.T(), loaded.Type, models.Savings.Slug)
}
