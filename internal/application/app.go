package application

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/gomoney/internal/application/ui"
	"github.com/PaulWaldo/gomoney/internal/db"
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

const (
	PrefKeyDBFile = "db_file"
)

type AppData struct {
	Service          *domain.Services
	Accounts         []models.Account
	Transactions     []models.Transaction
	DatabaseFileName binding.String
	// selectedAccount uint
	// UI Components
	accountList                     *widget.List
	transactionsTable               ui.TransactionsTable
	entryInfoPanel                  ui.EntryInfoPanel
	header                          ui.Header
	footer                          ui.Footer
	app                             fyne.App
	mainWindow                      fyne.Window
	accountAndTransactionsContainer *container.Split
	leftAndEntryInfo                *container.Split
	transactionEditRow              int
}

func (ad *AppData) accountSelected(id widget.ListItemID) {
	account := ad.Accounts[id]
	var err error
	ad.Transactions, err = ad.Service.Transaction.ListByAccount(account.ID)
	ad.transactionsTable.Table.Refresh()
	ad.footer.SetNumTransactions(len(ad.Transactions))

	if err != nil {
		panic(err)
	}
}

func (ad *AppData) onInfoButtonTapped() {
	ad.ToggleInfoPaneVisibility()
}

func (ad *AppData) HideInfoPane() {
	ad.entryInfoPanel.Form.Hide()
	ad.leftAndEntryInfo.SetOffset(1.0)
	ad.leftAndEntryInfo.Trailing.Refresh()
}

func (ad *AppData) UnhideInfoPane() {
	ad.entryInfoPanel.Form.Show()
	ad.leftAndEntryInfo.SetOffset(0.8)
	ad.leftAndEntryInfo.Trailing.Refresh()
}

func (ad *AppData) ToggleInfoPaneVisibility() {
	if ad.entryInfoPanel.Form.Visible() {
		ad.HideInfoPane()
	} else {
		ad.UnhideInfoPane()
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
	ad.header.InfoButton.OnTapped = ad.onInfoButtonTapped
	ad.footer = *ui.NewFooter()
	ad.accountList = ui.MakeAccountList(&ad.Accounts)
	ad.transactionsTable = ui.MakeTransactionsTable(&ad.Transactions, ad.mainWindow)
	ad.entryInfoPanel = *ui.MakeEntryInfoPanel()

	ad.accountList.OnSelected = ad.accountSelected
	ad.transactionsTable.SetOnSelectedCallback(ad.transactionSelected)
	ad.entryInfoPanel.Form.OnCancel = ad.onTransactionFormCancel
	ad.entryInfoPanel.Form.OnSubmit = ad.onTransactionFormSubmit
	ad.footer.SetNumTransactions(len(ad.Transactions))

	footerContainer := container.NewHBox(ad.footer.Label)

	ad.accountAndTransactionsContainer = container.NewHSplit(ad.accountList, ad.transactionsTable.Table)
	ad.accountAndTransactionsContainer.SetOffset(0.2)
	x := container.New(ui.NewCollapsibleLayout(), &ad.entryInfoPanel.Form)
	ad.leftAndEntryInfo = container.NewHSplit(ad.accountAndTransactionsContainer, x)
	ad.leftAndEntryInfo.SetOffset(0.8)
	ad.transactionsTable.Table.Refresh()

	return container.NewBorder(
		ad.header.Container, footerContainer, nil, nil,
		ad.leftAndEntryInfo,
	)
}

const useInMemoryDb = true

func (ad *AppData) openDatabase(file, migDir string) error {
	var err error
	if useInMemoryDb {
		ad.Service, err = db.NewSqliteInMemoryServices(migDir, true)
		if err != nil {
			return err
		}
	} else {
		dsn := fmt.Sprintf("file:%s?_foreign_keys=on", file)
		ad.Service, err = db.NewSqliteDiskServices(dsn, migDir)
		if err != nil {
			return err
		}
	}
	ad.Transactions, err = ad.Service.Transaction.List()
	if err != nil {
		return err
	}
	ad.Accounts, err = ad.Service.Account.List()
	if err != nil {
		return err
	}
	ad.app.Preferences().SetString(PrefKeyDBFile, file)
	ad.accountList.Refresh()
	ad.transactionsTable.Table.Refresh()
	ad.footer.SetNumTransactions(len(ad.Transactions))
	return nil
}

func (ad *AppData) chooseDatabaseFile(migDir string) (filename string, err error) {
	dialog.ShowFileOpen(func(rc fyne.URIReadCloser, e error) {
		if e != nil {
			dialog.ShowError(e, ad.mainWindow)
			err = e
			return
		}
		if rc == nil {
			return
		}
		filename := rc.URI().Path()
		fmt.Println(filename)
		err = ad.openDatabase(filename, migDir)
		if err != nil {
			dialog.ShowError(err, ad.mainWindow)
			return
		}
	}, ad.mainWindow)
	return filename, err
}

func (ad *AppData) createDatabaseFile(migDir string) {
	dialog.ShowFileSave(func(rc fyne.URIWriteCloser, e error) {
		if e != nil {
			dialog.ShowError(e, ad.mainWindow)
			return
		}
		if rc == nil {
			return
		}
		filename := rc.URI().Path()
		fmt.Println(filename)
		err := ad.openDatabase(filename, migDir)
		if err != nil {
			dialog.ShowError(err, ad.mainWindow)
			return
		}
	}, ad.mainWindow)
}

func (ad *AppData) loadDefaults(migDir string) {
	dbFile := ad.app.Preferences().String(PrefKeyDBFile)
	if len(dbFile) == 0 {
		return
	}
	// Stat the file first, as Sqlite will happily "open" a filename that does not exist
	if _, err := os.Stat(dbFile); err != nil {
		e := fmt.Errorf("unable to open database \"%s\"\n%w", dbFile, err)
		dialog.ShowError(e, ad.mainWindow)
		return
	}
	if err := ad.openDatabase(dbFile, migDir); err != nil {
		e := fmt.Errorf("unable to open database \"%s\"\n%w", dbFile, err)
		dialog.ShowError(e, ad.mainWindow)
		return
	}
}

func (ad *AppData) setupWindowTitleListener() {
	dbFileBinding := binding.BindPreferenceString(PrefKeyDBFile, ad.app.Preferences())
	titleBinding := binding.StringToStringWithFormat(dbFileBinding, "MoneyMinder - %s")
	l := binding.NewDataListener(func() {
		t, _ := titleBinding.Get()
		ad.mainWindow.SetTitle(t)
	})
	titleBinding.AddListener(l)
}

const appId = "com.github.paulwaldo.gomoney"

func RunApp(ad *AppData) {
	ad.app = app.NewWithID(appId)
	ad.mainWindow = ad.app.NewWindow("MoneyMinder")
	ad.setupWindowTitleListener()
	const migrationsDir = "../../internal/db/migrations"
	ad.mainWindow.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New...", func() { ad.createDatabaseFile(migrationsDir) }),
			fyne.NewMenuItem("Open...", func() { ad.chooseDatabaseFile(migrationsDir) }),
		)),
	)
	ad.mainWindow.Resize(fyne.NewSize(1000, 600))
	ad.mainWindow.SetContent(ad.makeUI(ad.mainWindow))
	ad.loadDefaults(migrationsDir)

	ad.mainWindow.ShowAndRun()
}
