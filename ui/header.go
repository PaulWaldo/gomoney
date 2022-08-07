package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ad AppData) makeHeader() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel("MoneyMinder"))
}
