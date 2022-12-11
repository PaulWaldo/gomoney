package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/gomoney/data"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

type AccountsPanel struct {
	/*Container*/ *fyne.Container
	List              *widget.List
	AddButton         *widget.Button
	EditButton        *widget.Button
	SelectedAccountId widget.ListItemID
	accounts          *[]models.Account
	parent            *fyne.Window
}

func (ap *AccountsPanel) makeBottom(newAccountCallback func(bool, models.Account), editAccountCallback func(bool, models.Account)) *fyne.Container {
	ap.AddButton = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		ap.NewAccount(*ap.parent, newAccountCallback)
	})
	ap.EditButton = widget.NewButtonWithIcon("", data.IconCreateblack24dpSvg, func() {
		if ap.SelectedAccountId == 0 {
			fmt.Println("Cannot edit All Accounts account")
			return
		}
		selectedAcct := (*ap.accounts)[ap.SelectedAccountId-1]
		ap.EditAccount(*ap.parent, selectedAcct, editAccountCallback)
	})
	return container.NewHBox(ap.AddButton, ap.EditButton)
}

func MakeAccountsPanel(
	accounts *[]models.Account,
	parent *fyne.Window,
	newAccountCallback func(bool, models.Account),
	editAccountCallback func(bool, models.Account)) *AccountsPanel {

	ap := AccountsPanel{}
	ap.accounts = accounts
	ap.parent = parent
	ap.List = widget.NewList(
		// length
		func() int {
			return len(*accounts) + 1
		},
		// createItem
		func() fyne.CanvasObject {
			return widget.NewLabel("My Checking Account")
		},
		// updateItem
		func(i widget.ListItemID, o fyne.CanvasObject) {
			var text string
			if i == 0 {
				text = "All Accounts"
			} else {
				text = (*accounts)[i-1].Name
			}
			o.(*widget.Label).SetText(text)
		},
	)
	bottom := ap.makeBottom(newAccountCallback, editAccountCallback)
	ap.Container = container.NewBorder(nil, bottom, nil, nil, ap.List)
	return &ap
}

func (ap AccountsPanel) makeForm(title string, confirm string, dismiss string, account *models.Account, callback func(bool), parent fyne.Window) dialog.Dialog {
	name := widget.NewEntryWithData(binding.BindString(&account.Name))
	name.SetPlaceHolder("My checking")
	name.Validator = validation.NewRegexp(`.+`, "name required")

	acctType := widget.NewSelect(
		[]string{models.Checking.Slug, models.Savings.Slug, models.CreditCard.Slug},
		func(val string) { account.Type = val })
	acctType.SetSelectedIndex(0)

	memo := widget.NewEntryWithData(binding.BindString(&account.Memo))
	memo.MultiLine = true
	memo.SetPlaceHolder("Free-form\ntext")

	routing := widget.NewEntryWithData(binding.BindString(&account.Routing))
	routing.SetPlaceHolder("123456789")
	routing.Validator = validation.NewRegexp(`^$|\d{9}`, "invalid routing number")

	acctNumber := widget.NewEntryWithData(binding.BindString(&account.AccountNumber))
	acctNumber.SetPlaceHolder("Account #")

	hidden := widget.NewCheckWithData("Hide", binding.BindBool(&account.Hidden))
	netWorthInclude := widget.NewCheckWithData("Include", binding.BindBool(&account.NetWorthInclude))
	budgetInclude := widget.NewCheckWithData("Include", binding.BindBool(&account.NetWorthInclude))

	items := []*widget.FormItem{
		{Text: "Name", Widget: name, HintText: "Your full name"},
		{Text: "Type", Widget: acctType, HintText: "Account Type"},
		{Text: "Memo", Widget: memo, HintText: "Free-form text"},
		{Text: "Routing #", Widget: routing, HintText: "Routing Number"},
		{Text: "Number", Widget: acctNumber, HintText: "Account Number"},
		{Text: "Hidden", Widget: hidden, HintText: "Hidden"},
		{Text: "Net Worth", Widget: netWorthInclude, HintText: "Net Worth Include"},
		{Text: "Budget", Widget: budgetInclude, HintText: ""},
		// {Text: "", Widget: , HintText: ""},
	}
	d := dialog.NewForm(title, confirm, dismiss, items, callback, parent)
	return d
}

func (ap AccountsPanel) NewAccount(parent fyne.Window, callback func(bool, models.Account)) {
	newAcct := models.Account{}
	form := ap.makeForm("New Account", "Create", "Cancel", &newAcct, func(b bool) {
		callback(b, newAcct)
	}, parent)
	form.Show()
}

func (ap AccountsPanel) EditAccount(parent fyne.Window, account models.Account, callback func(bool, models.Account)) {
	form := ap.makeForm("Edit Account", "Create", "Cancel", &account, func(b bool) {
		callback(b, account)
	}, parent)
	form.Show()
}
