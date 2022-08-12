package ui

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

type Footer struct {
	Label *widget.Label
}

func NewFooter() Footer {
	return Footer{Label: widget.NewLabel("Footer")}
}

func (f Footer) SetNumTransactions(n int64) {
	f.Label.SetText(fmt.Sprintf("%d Transactions", n))
}
