package ui

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

type Footer struct {
	Label *widget.Label
	// Container *fyne.Container
}

func NewFooter() *Footer {
	w := widget.NewLabel("Footer")
	return &Footer{Label: w /*Container: container.NewHBox(w)*/}
}

func (f *Footer) SetNumTransactions(n int64) {
	f.Label.SetText(fmt.Sprintf("%d Transactions", n))
}
