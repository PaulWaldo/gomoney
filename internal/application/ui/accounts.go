package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

func MakeAccountList(accounts *[]models.Account) *widget.List {
	return widget.NewList(
		// length
		func() int {
			return len(*accounts)
		},
		// createItem
		func() fyne.CanvasObject {
			return widget.NewLabel("My Checking Account")
		},
		// updateItem
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText((*accounts)[i].Name)
		},
	)
}
