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
			return len(*accounts) + 1
		},
		// createItem
		func() fyne.CanvasObject {
			return widget.NewLabel("My Checking Account")
		},
		// updateItem
		func(i widget.ListItemID, o fyne.CanvasObject) {
			var text string
			if i == 0 {
				text = "All Accounts"
			} else {
				text = (*accounts)[i-1].Name
			}
			o.(*widget.Label).SetText(text)
		},
	)
}
