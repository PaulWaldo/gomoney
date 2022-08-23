package application

import (
	"fmt"

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
	accountList        *widget.List
	transactionsTable  ui.TransactionsTable
	entryInfoPanel     ui.EntryInfoPanel
	header             ui.Header
	footer             ui.Footer
	app                fyne.App
	mainWindow         fyne.Window
	transactionEditRow int
}

func (ad *AppData) accountSelected(id widget.ListItemID) {
	account := ad.Accounts[id]
	var err error
	var count int64
	ad.Transactions, count, err = ad.Service.Transaction.ListByAccount(account.ID)
	// ad.transactionsTable.Refresh()
	ad.footer.SetNumTransactions(count)

	if err != nil {
		panic(err)
	}
}

func (ad *AppData) transactionSelected(i widget.TableCellID) {
	ad.transactionEditRow = i.Row
	ad.entryInfoPanel.SetTransaction(ad.Transactions[i.Row])
}

func (ad *AppData) onTransactionFormSubmit() {
	fmt.Println("Submitting")
	ad.entryInfoPanel.OnSubmit()
	ad.Transactions[ad.transactionEditRow] = ad.entryInfoPanel.Transaction
	ad.transactionsTable.Table.Refresh()
}

func (ad *AppData) onTransactionFormCancel() {
	ad.entryInfoPanel.OnCancel()
	fmt.Println("Cancelling")
}

func (ad *AppData) makeUI(mainWindow fyne.Window) *fyne.Container {
	// ad.SetSelectedAccount(0)
	ad.header = ui.MakeHeader()
	ad.footer = *ui.NewFooter()
	ad.accountList = ui.MakeAccountList(&ad.Accounts)
	ad.transactionsTable = ui.MakeTransactionsTable(&ad.Transactions, ad.mainWindow)
	ad.entryInfoPanel = *ui.MakeEntryInfoPanel()

	ad.accountList.OnSelected = ad.accountSelected
	ad.transactionsTable.SetOnSelectedCallback(ad.transactionSelected)
	ad.entryInfoPanel.Form.OnCancel = ad.onTransactionFormCancel
	ad.entryInfoPanel.Form.OnSubmit = ad.onTransactionFormSubmit
	ad.footer.SetNumTransactions(int64(len(ad.Transactions)))

	footerContainer := container.NewHBox(ad.footer.Label)

	accountsAndTransactions := container.NewHSplit(ad.accountList, ad.transactionsTable.Table)
	accountsAndTransactions.SetOffset(0.2)
	allSplits := container.NewHSplit(accountsAndTransactions, &ad.entryInfoPanel.Form)
	allSplits.SetOffset(0.8)
	ad.transactionsTable.Table.Refresh()

	return container.NewBorder(
		ad.header.Container, footerContainer, nil, nil,
		allSplits,
	)
}

func RunApp(ad *AppData) {
	ad.app = app.New()
	ad.mainWindow = ad.app.NewWindow("MoneyMinder")
	ad.mainWindow.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File", fyne.NewMenuItem("Open...", func() {})),
	))
	ad.mainWindow.Resize(fyne.NewSize(1000, 600))
	ad.mainWindow.SetContent(ad.makeUI(ad.mainWindow))
	ad.header.InfoButton.OnTapped = ad.modifyTransaction
	ad.mainWindow.ShowAndRun()
}

func (ad *AppData) modifyTransaction() {
}
