package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const numAccounts = 10
const numTransactions = 50

var accts []string

type transaction struct {
	Name   string
	Amount float64
	Date   time.Time
	Memo   string
}

var transactions []transaction

func init() {
	accts = make([]string, numAccounts)
	for i := range accts {
		accts[i] = fmt.Sprintf("Account %d", i)
	}

	transactions = make([]transaction, numTransactions)
	amt := 1.0
	for i := range transactions {
		transactions[i].Amount = amt + float64(i)
		transactions[i].Name = fmt.Sprintf("Trans %d", i)
		transactions[i].Date = time.Now().Add(-time.Hour * time.Duration(i))
		transactions[i].Memo = fmt.Sprintf("Memo for %d", i)
	}
}

func makeHeader() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel("Header"))
}

func makeLeftSidebar() *fyne.Container {
	list := widget.NewList(
		func() int {
			return /*rand.Intn*/ (numAccounts)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(accts[i])
		},
	)
	return container.NewMax(list)
}

func makeCenter() *fyne.Container {
	list := widget.NewTable(
		func() (int, int) {
			return len(transactions), 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			switch i.Col {
			case 0:
				o.(*widget.Label).SetText(transactions[i.Row].Name)
			case 1:
				o.(*widget.Label).SetText(transactions[i.Row].Date.String())
			case 2:
				o.(*widget.Label).SetText(fmt.Sprintf("%.2f", transactions[i.Row].Amount))
			case 3:
				o.(*widget.Label).SetText(transactions[i.Row].Memo)
			}
		})
	return container.NewMax(list)
}

func makeUI() *fyne.Container {
	header := makeHeader()
	leftSideBar := makeLeftSidebar()
	return container.NewBorder(header, nil, leftSideBar, nil, header, leftSideBar, makeCenter())
}

func RunApp() {
	a := app.New()
	w := a.NewWindow("MoneyMinder")
	w.SetContent(makeUI())
	w.ShowAndRun()
}
