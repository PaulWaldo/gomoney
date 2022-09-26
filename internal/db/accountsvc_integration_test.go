package db

import (
	"testing"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const migrationDir = "migrations"

type accountServiceTestSuite struct {
	suite.Suite
	Services *domain.Services
	AcctSvc  domain.AccountSvc
}

func (suite *accountServiceTestSuite) SetupTest() {
	services, err := NewSqliteInMemoryServices(migrationDir, false)
	require.NoError(suite.T(), err)
	suite.AcctSvc = services.Account
	suite.Services = services
}

func (suite *accountServiceTestSuite) TearDownTest() {
	suite.Services.Db.Close()
}

func TestAccountServiceTestSuite(t *testing.T) {
	suite.Run(t, new(accountServiceTestSuite))
}

func (suite *accountServiceTestSuite) Test_accountSvc_Create() {
	type fields struct {
	}
	type args struct {
		account models.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		{
			name:    "create success returns id",
			fields:  fields{},
			args:    args{account: models.Account{Name: "My Checking", Type: models.Checking.Slug, Memo: "m1", Routing: "123", AccountNumber: "567", Hidden: false, NetWorthInclude: false, BudgetInclude: false}},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			err := suite.AcctSvc.Create(&tt.args.account)
			if (err != nil) != tt.wantErr {
				suite.T().Fatalf("accountSvc.Create() error = '%v', wantErr %v", err, tt.wantErr)
				return
			}

			got, err := suite.AcctSvc.Get(tt.args.account.ID)
			assert.NoError(suite.T(), err)
			assert.EqualValues(suite.T(), got, tt.args.account)
		})
	}
}

func (suite *accountServiceTestSuite) Test_accountSvc_Update() {
	saved := models.Account{
		Name: "Account",
		Type: models.Checking.Slug,
	}

	err := suite.AcctSvc.Create(&saved)
	require.NoError(suite.T(), err)

	saved.Name = "Changed"
	saved.Type = models.Savings.Slug
	err = suite.AcctSvc.Update(&saved)
	require.NoError(suite.T(), err)

	loaded, err := suite.AcctSvc.Get(saved.ID)
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), loaded.Name, "Changed")
	assert.Equal(suite.T(), loaded.Type, models.Savings.Slug)
}

func (suite *accountServiceTestSuite) Test_accountSvc_Get() {
	saved := models.Account{
		Name:            "Account",
		Type:            models.Checking.Slug,
		Memo:            "memo",
		Routing:         "1234567",
		AccountNumber:   "246810",
		Hidden:          true,
		NetWorthInclude: true,
		BudgetInclude:   false,
	}

	err := suite.AcctSvc.Create(&saved)
	require.NoError(suite.T(), err)

	loaded, err := suite.AcctSvc.Get(saved.ID)
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), loaded.Name, "Account")
	assert.Equal(suite.T(), loaded.Type, models.Checking.Slug)
}

func (suite *accountServiceTestSuite) Test_accountSvc_List() {
	accounts := []models.Account{
		{Name: "a1", Type: models.Checking.Slug, Memo: "m1", Routing: "123", AccountNumber: "567", Hidden: false, NetWorthInclude: false, BudgetInclude: false},
		{Name: "a2", Type: models.Savings.Slug, Memo: "m2", Routing: "123", AccountNumber: "567", Hidden: true, NetWorthInclude: true, BudgetInclude: true},
		{Name: "a3", Type: models.CreditCard.Slug, Memo: "m3", Routing: "123", AccountNumber: "567", Hidden: false, NetWorthInclude: false, BudgetInclude: false},
	}

	saved := []models.Account{}
	for _, a := range accounts {
		err := suite.AcctSvc.Create(&a)
		require.NoError(suite.T(), err)
		saved = append(saved, a)
	}

	loaded, err := suite.AcctSvc.List()
	require.NoError(suite.T(), err)
	assert.EqualValues(suite.T(), len(accounts), len(loaded))
	for i := range loaded {
		want := saved[i]
		got := loaded[i]
		assert.Equal(suite.T(), want, got)
		assert.Equal(suite.T(), want.ID, got.ID)
		assert.Equal(suite.T(), want.Name, got.Name)
		assert.Equal(suite.T(), want.Type, got.Type)
	}
}
