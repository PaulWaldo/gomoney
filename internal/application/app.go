package application

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/gomoney/internal/application/ui"
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

type AppData struct {
	Service         domain.Services
	Accounts        []models.Account
	Transactions    []models.Transaction
	selectedAccount uint
	// UI Components
	accountList       *widget.List
	transactionsTable *widget.Table
	header            *fyne.Container
	footer            *fyne.Container
}

func (ad *AppData) accountSelected(id widget.ListItemID) {
	account := ad.Accounts[id]
	var err error
	ad.Transactions, _, err = ad.Service.Transaction.ListByAccount(account.ID)
	if err != nil {
		panic(err)
	}
}

// func (ad AppData) makeLeftSidebar() *fyne.Container {
// 	ad.accountList = NewAccountList(&ad.Accounts)
// 	ad.accountList.widget.OnSelected = ad.accountSelected
// 	return container.NewMax(ad.accountList.widget)
// }

// func (ad AppData) makeAccountList() *widget.List {
// 	ad.accountList = NewAccountList(&ad.Accounts)
// 	ad.accountList.widget.OnSelected = ad.accountSelected
// 	return ad.accountList.widget
// }

// func (ad AppData) makeTransactionsTable() *widget.Table{
// 	table := NewTransactionsTable(&ad.Transactions)
// 	return table.table
// }

func (ad *AppData) makeUI() *fyne.Container {
	// ad.SetSelectedAccount(0)
	ad.header = ui.MakeHeader()
	ad.footer = ui.MakeFooter()
	ad.accountList = ui.MakeAccountList(&ad.Accounts)
	ad.accountList.OnSelected = ad.accountSelected
	ad.transactionsTable = ui.MakeTransactionsTable(&ad.Transactions)

	center := container.NewHSplit(ad.accountList, ad.transactionsTable)
	center.Offset = 0.2

	return container.NewBorder(ad.header, ad.footer, nil, nil /*header, footer,*/, center)
}

func RunApp(ad *AppData) {
	a := app.New()
	w := a.NewWindow("MoneyMinder")
	w.Resize(fyne.NewSize(600, 400))
	w.SetContent(ad.makeUI())
	w.ShowAndRun()
}
