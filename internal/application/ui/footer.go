package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Footer struct {
	Label *widget.Label
}

func (f *Footer) MakeFooter() *fyne.Container {
	f.Label = widget.NewLabel("Footer")
	return container.NewHBox(f.Label)
}
