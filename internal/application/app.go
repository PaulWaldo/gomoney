package application

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/gomoney/internal/application/ui"
	"github.com/PaulWaldo/gomoney/internal/db"
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

const (
	PrefKeyDBFile           = "db_file"
	PrefKeyLastOpenLocation = "last_open_dir"
)

type AppData struct {
	Service          *domain.Services
	Accounts         []models.Account
	Transactions     []models.Transaction
	DatabaseFileName binding.String
	// selectedAccount uint
	// UI Components
	accountsPanel                   *ui.AccountsPanel
	transactionsTable               ui.TransactionsTable
	entryInfoPanel                  ui.EntryInfoPanel
	header                          ui.Header
	footer                          ui.Footer
	app                             fyne.App
	mainWindow                      fyne.Window
	accountAndTransactionsContainer *container.Split
	leftAndEntryInfo                *container.Split
	transactionEditRow              int
	LoadSampleData                  bool
	InMemDatabase                   bool
}

func (ad *AppData) onAccountSelected(id widget.ListItemID) {
	var err error
	if id == 0 {
		ad.Transactions, err = ad.Service.Transaction.List()
	} else {
		account := ad.Accounts[id-1]
		ad.Transactions, err = ad.Service.Transaction.ListByAccount(account.ID)
	}
	ad.transactionsTable.Table.Refresh()
	ad.accountsPanel.SelectedAccountId = id
	ad.footer.SetNumTransactions(len(ad.Transactions))

	if err != nil {
		panic(err)
	}
}

func (ad *AppData) onInfoButtonTapped() {
	ad.ToggleInfoPaneVisibility()
}

func (ad *AppData) onTransactionAddButtonTapped() {
	newTx := models.Transaction{Date: time.Now(), AccountID: int64(ad.accountsPanel.SelectedAccountId)}
	ad.Transactions = append(ad.Transactions, newTx)
	ad.Service.Transaction.Create(&newTx)
	ad.transactionsTable.Table.Refresh()
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

func (ad *AppData) onNewAccount(saved bool, editedAccount models.Account) {
	if saved {
		err := ad.Service.Account.Create(&editedAccount)
		if err != nil {
			dialog.NewError(err, ad.mainWindow)
			return
		}

		// Create an "Initial Balance Record"
		ib := models.Transaction{
			Date: time.Now(),
			Payee: "Initial Balance",
			AccountID: editedAccount.ID}
		err = ad.Service.Transaction.Create(&ib)
		if err != nil {
			dialog.NewError(err, ad.mainWindow)
			return
		}

		ad.Accounts, err = ad.Service.Account.List()
		if err != nil {
			dialog.NewError(err, ad.mainWindow)
			return
		}
		ad.accountsPanel.List.Refresh()
		ad.accountsPanel.List.Select(1)
	}
}

func (ad *AppData) onEditAccount(saved bool, editedAccount models.Account) {
	if saved {
		err := ad.Service.Account.Update(&editedAccount)
		if err != nil {
			dialog.NewError(err, ad.mainWindow)
			return
		}
		ad.Accounts, err = ad.Service.Account.List()
		if err != nil {
			dialog.NewError(err, ad.mainWindow)
			return
		}
		ad.accountsPanel.List.Refresh()
	}
}

func (ad *AppData) makeUI(mainWindow fyne.Window) *fyne.Container {
	// ad.SetSelectedAccount(0)
	ad.header = ui.MakeHeader()
	ad.header.InfoButton.OnTapped = ad.onInfoButtonTapped
	ad.header.AddButton.OnTapped = ad.onTransactionAddButtonTapped
	ad.footer = *ui.NewFooter()
	ad.accountsPanel = ui.MakeAccountsPanel(&ad.Accounts, &mainWindow, ad.onNewAccount, ad.onEditAccount)
	ad.transactionsTable = ui.MakeTransactionsTable(&ad.Transactions, ad.mainWindow)
	ad.entryInfoPanel = *ui.MakeEntryInfoPanel()

	ad.accountsPanel.List.OnSelected = ad.onAccountSelected
	ad.transactionsTable.SetOnSelectedCallback(ad.transactionSelected)
	ad.entryInfoPanel.Form.OnCancel = ad.onTransactionFormCancel
	ad.entryInfoPanel.Form.OnSubmit = ad.onTransactionFormSubmit
	ad.footer.SetNumTransactions(len(ad.Transactions))

	footerContainer := container.NewHBox(ad.footer.Label)

	ad.accountAndTransactionsContainer = container.NewHSplit(ad.accountsPanel.Container, ad.transactionsTable.Table)
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

func (ad *AppData) openDatabase(file, migDir string) error {
	var err error
	if ad.InMemDatabase {
		ad.Service, err = db.NewSqliteInMemoryServices(migDir, ad.LoadSampleData)
		if err != nil {
			return err
		}
	} else {
		dsn := fmt.Sprintf("file:%s", file)
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
	ad.accountsPanel.List.Refresh()
	ad.transactionsTable.Table.Refresh()
	ad.footer.SetNumTransactions(len(ad.Transactions))
	return nil
}

func (ad *AppData) chooseDatabaseFile(migDir string) (filename string, err error) {
	var chosenURI fyne.URI
	d := dialog.NewFileOpen(func(rc fyne.URIReadCloser, e error) {
		if e != nil {
			dialog.ShowError(e, ad.mainWindow)
			err = e
			return
		}
		if rc == nil {
			return
		}
		chosenURI = rc.URI()
		fmt.Printf("Chosen URI: %s\n", chosenURI)
		filename := rc.URI().Path()
		fmt.Println(filename)
		err = ad.openDatabase(filename, migDir)
		if err != nil {
			dialog.ShowError(err, ad.mainWindow)
			return
		}
		ad.app.Preferences().SetString(PrefKeyLastOpenLocation, chosenURI.String())
	}, ad.mainWindow)

	lastLoc := ad.app.Preferences().String(PrefKeyLastOpenLocation)
	if len(lastLoc) > 0 {
		// Use stored location
		lastURI, err := storage.ParseURI(lastLoc)
		if err != nil {
			return "", err
		}
		lastLocDir, err := storage.Parent(lastURI)
		if err != nil {
			return "", err
		}
		listableURI, err := storage.ListerForURI(lastLocDir)
		if err != nil {
			return "", err
		}
		d.SetLocation(listableURI)
	}
	d.Show()
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
	const migrationsDir = "internal/db/migrations"
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
