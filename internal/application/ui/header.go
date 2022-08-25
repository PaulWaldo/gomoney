package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Header struct {
	AppNameLabel *widget.Label
	InfoButton   *widget.Button
	AddButton    *widget.Button
	Container    *fyne.Container
}

func MakeHeader() Header {
	h := Header{}
	h.AppNameLabel = widget.NewLabel("MoneyMinder")
	h.InfoButton = widget.NewButtonWithIcon("", theme.InfoIcon(), func() {})
	h.AddButton = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {})
	h.Container = container.NewBorder(nil, nil, h.AppNameLabel, container.NewHBox(h.InfoButton, h.AddButton))
	return h
}
