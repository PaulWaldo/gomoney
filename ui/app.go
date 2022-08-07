package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

const YYYYMMDD = "2006-01-02"

type AppData struct {
	Accounts     []models.Account
	Transactions []models.Transaction
}

func (ma AppData) makeHeader() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel("MoneyMinder"))
}

func (ma AppData) makeFooter() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel("Footer"))
}

func (ma AppData) makeLeftSidebar() *fyne.Container {
	var bindings []binding.DataMap
	for i := range ma.Accounts{
		bindings = append(bindings, binding.BindStruct(&ma.Accounts[i]))
	}
	list := widget.NewList(
		func() int {
			return len(ma.Accounts)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("My Checking Account")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			name, err := bindings[i].GetItem("Name")
			if err != nil {
				panic(err)
			}
			x,err:=name.(binding.String).Get()
			if err != nil {
				panic(err)
			}
			o.(*widget.Label).SetText(x)
		},
	)
	return container.NewMax(list)
}

func (ma AppData) makeCenter() *fyne.Container {
	var bindings []binding.DataMap
	for i := range ma.Transactions{
		bindings = append(bindings, binding.BindStruct(&ma.Transactions[i]))
	}
	table := widget.NewTable(
		func() (int, int) {
			return len(ma.Transactions), 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			switch i.Col {
			case 0:
				payee, err := bindings[i.Row].GetItem("Payee")
				if err != nil {
					panic(err)
				}
				x,err:=payee.(binding.String).Get()
				if err != nil {
					panic(err)
				}
				o.(*widget.Label).SetText(x)
			case 1:
				o.(*widget.Label).SetText(ma.Transactions[i.Row].Date.Format(YYYYMMDD))
			case 2:
				o.(*widget.Label).SetText(fmt.Sprintf("%.2f", ma.Transactions[i.Row].Amount))
			case 3:
				o.(*widget.Label).SetText(ma.Transactions[i.Row].Memo)
			}
		},
	)
	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 300)
	split := container.NewHSplit(ma.makeLeftSidebar(), table)
	split.Offset = 0.2

	return container.NewMax(split)
}

func (ma AppData) makeUI() *fyne.Container {
	header := ma.makeHeader()
	footer := ma.makeFooter()
	return container.NewBorder(header, footer, nil, nil, header, footer, ma.makeCenter())
}

func (ma AppData) RunApp() {
	a := app.New()
	w := a.NewWindow("MoneyMinder")
	w.Resize(fyne.NewSize(600, 400))
	w.SetContent(ma.makeUI())
	w.ShowAndRun()
}
