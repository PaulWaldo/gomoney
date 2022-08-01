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
	Name string
	Amount float64
	Date time.Time
	Memo string
}
var transactions []transaction

func init() {
	accts = make([]string, numAccounts)
	for i := range accts {
		accts[i] = fmt.Sprintf("Account %d", i)
	}

	transactions = make([]transaction, numTransactions)
	amt := 1.0
	for i, t := range transactions {
		t.Amount = amt+float64(i)
		t.Name = fmt.Sprintf("Trans %d", i)
		t.Date=time.Now().Add
	}
}

func makeHeader() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel("Header"))
}

func makeLeftSidebar() *fyne.Container {
	list := widget.NewList(
		func() int {
			return /*rand.Intn*/(numAccounts)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(accts[i])
		},
	)
	return container.NewMax(list)
	// const numAccounts = 10
	// accts := make([]fyne.CanvasObject, numAccounts)
	// for i := range accts {
	// 	accts[i] = widget.NewLabel(fmt.Sprintf("Account %d", i))
	// }
	// return container.NewVBox(accts...)
}

func makeCenter() *fyne.Container {

}

func makeUI() *fyne.Container {
	header := makeHeader()
	leftSideBar := makeLeftSidebar()
	return container.NewBorder(header, nil, leftSideBar, nil, header, leftSideBar, widget.NewLabel("Center"))
}

func RunApp() {
	a := app.New()
	w := a.NewWindow("MoneyMinder")
	w.SetContent(makeUI())
	w.ShowAndRun()
}
