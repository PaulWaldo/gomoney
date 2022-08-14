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
	Service      domain.Services
	Accounts     []models.Account
	Transactions []models.Transaction
	// selectedAccount uint
	// UI Components
	accountList       *widget.List
	transactionsTable *widget.Table
	header            ui.Header
	footer            ui.Footer
}

func (ad *AppData) accountSelected(id widget.ListItemID) {
	account := ad.Accounts[id]
	var err error
	var count int64
	ad.Transactions, count, err = ad.Service.Transaction.ListByAccount(account.ID)
	ad.transactionsTable.Refresh()
	ad.footer.SetNumTransactions(count)

	if err != nil {
		panic(err)
	}
}

func (ad *AppData) makeUI() *fyne.Container {
	// ad.SetSelectedAccount(0)
	ad.header = ui.MakeHeader()
	ad.footer = *ui.NewFooter()
	footer := container.NewHBox(ad.footer.Label)
	ad.accountList = ui.MakeAccountList(&ad.Accounts)
	ad.accountList.OnSelected = ad.accountSelected
	ad.transactionsTable = ui.MakeTransactionsTable(&ad.Transactions)
	ad.footer.SetNumTransactions(int64(len(ad.Transactions)))

	center := container.NewHSplit(ad.accountList, ad.transactionsTable)
	center.Offset = 0.2

	return container.NewBorder(ad.header.Container, footer, nil, nil /*header, footer,*/, center)
}

func RunApp(ad *AppData) {
	a := app.New()
	w := a.NewWindow("MoneyMinder")
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File", fyne.NewMenuItem("Open...", func() {})),
	))
	w.Resize(fyne.NewSize(1000, 600))
	w.SetContent(ad.makeUI())
	// ad.transactionsTable.Refresh()
	w.ShowAndRun()
}
