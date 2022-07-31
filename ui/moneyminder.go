package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeHeader() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel("Header"))
}

func makeLeftSidebar() *fyne.Container {
	const numAccounts = 10
	accts := make([]*widget.Label, numAccounts)
	for i := range accts {
		accts[i] = widget.NewLabel(fmt.Sprintf("Accountd %d", i))
	}
	return container.NewVBox(accts...)//[0], accts[1], accts[2])
}

func makeUI() *fyne.Container {
	header := makeHeader()
	leftSideBar := makeLeftSidebar()
	return container.NewBorder(header, nil, leftSideBar, nil, header, leftSideBar)
	// return widget.NewLabel("Hello world!"),
	// 	widget.NewEntry()
}

func RunApp() {
	a := app.New()
	w := a.NewWindow("MoneyMinder")
	w.SetContent(makeUI())
	w.ShowAndRun()
}
