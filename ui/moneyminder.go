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
const YYYYMMDD = "2006-01-02"

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
		transactions[i].Date = time.Now().Add(-time.Hour * 24 * time.Duration(i))
		transactions[i].Memo = fmt.Sprintf("Memo for %d", i)
	}
}

func makeHeader() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel("Header"))
}

func makeFooter() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel("Footer"))
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
	table := widget.NewTable(
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
				o.(*widget.Label).SetText(transactions[i.Row].Date.Format(YYYYMMDD))
			case 2:
				o.(*widget.Label).SetText(fmt.Sprintf("%.2f", transactions[i.Row].Amount))
			case 3:
				o.(*widget.Label).SetText(transactions[i.Row].Memo)
			}
		},
	)
	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 300)
	split := container.NewHSplit(makeLeftSidebar(), table)
	split.Offset = 0.2

	return container.NewMax(split)

}

func makeUI() *fyne.Container {
	header := makeHeader()
	// leftSideBar := makeLeftSidebar()
	footer := makeFooter()
	return container.NewBorder(header, footer, nil, nil, header, footer, makeCenter())
}

func RunApp() {
	a := app.New()
	w := a.NewWindow("MoneyMinder")
	w.SetContent(makeUI())
	w.ShowAndRun()
}
