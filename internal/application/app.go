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
	accountList       *widget.List
	transactionsTable ui.TransactionsTable
	entryInfoPanel    ui.EntryInfoPanel
	header            ui.Header
	footer            ui.Footer
	app               fyne.App
	mainWindow        fyne.Window
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

func (ad *AppData) makeUI(mainWindow fyne.Window) *fyne.Container {
	// ad.SetSelectedAccount(0)
	ad.header = ui.MakeHeader()
	ad.footer = *ui.NewFooter()
	ad.accountList = ui.MakeAccountList(&ad.Accounts)
	ad.transactionsTable = ui.MakeTransactionsTable(&ad.Transactions, ad.mainWindow)
	ad.entryInfoPanel = *ui.MakeEntryInfoPanel()

	ad.accountList.OnSelected = ad.accountSelected
	ad.footer.SetNumTransactions(int64(len(ad.Transactions)))

	// coloredRect := canvas.NewRectangle(color.RGBA{R: 128, A: 128})
	accountsAndTransactions:=container.NewHSplit(ad.accountList, ad.transactionsTable.Table)
	allSplits:= container.NewHSplit(accountsAndTransactions, &ad.entryInfoPanel.Form)
	center := container.NewHSplit(
		ad.accountList,
		container.NewBorder(nil, nil, nil, &ad.entryInfoPanel.Form,
			ad.transactionsTable.Table,
		),
	)
	// ad.entryInfoPanel.Form.Hide()
	fmt.Printf("Table MinSize: %v\n", ad.transactionsTable.Table.MinSize())
	fmt.Printf("InfoPanel MinSize: %v\n", ad.entryInfoPanel.Form.MinSize())
	center.Offset = 0.2

	return container.NewBorder(
		ad.header.Container, ad.footer.Container, nil, nil,
		center,
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
	// ad.transactionsTable.Refresh()
	ad.mainWindow.ShowAndRun()
}

func (ad *AppData) modifyTransaction() {
	i := ui.InfoFormDialog{Parent: ad.mainWindow}
	i.ShowInfoForm()
}
