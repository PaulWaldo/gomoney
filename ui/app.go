package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

const YYYYMMDD = "2006-01-02"

type AppData struct {
	Accounts     []models.Account
	Transactions []models.Transaction
}

func (ad AppData) makeUI() *fyne.Container {
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
