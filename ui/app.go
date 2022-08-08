package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

const YYYYMMDD = "2006-01-02"

type AppData struct {
	Service domain.Services
	Accounts     []models.Account
	Transactions []models.Transaction
	selectedAccount uint
}

func (ad AppData) makeUI() *fyne.Container {
	ad.SetSelectedAccount(0)
	header := ad.makeHeader()
	footer := ad.makeFooter()
	return container.NewBorder(header, footer, nil, nil, header, footer, ad.makeCenter())
}

func (ad AppData) RunApp() {
	a := app.New()
	w := a.NewWindow("MoneyMinder")
	w.Resize(fyne.NewSize(600, 400))
	w.SetContent(ad.makeUI())
	w.ShowAndRun()
}
