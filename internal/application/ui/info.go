package ui

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

type EntryInfoPanel struct {
	Form        widget.Form
	payee       *widget.Entry
	amount      *widget.Entry
	memo        *widget.Entry
	date        *widget.Entry
	transaction *models.Transaction
}

// func (eip *EntryInfoPanel) onSubmit() {
// 	fmt.Println("Submitting")
// }

// func (eip *EntryInfoPanel) onCancel() {
// 	fmt.Println("Cancelling")
// }

func MakeEntryInfoPanel() *EntryInfoPanel {
	eip := &EntryInfoPanel{
		payee:  widget.NewEntry(),
		amount: widget.NewEntry(),
		memo:   widget.NewMultiLineEntry(),
		date:   widget.NewEntry(),
	}
	eip.Form = *widget.NewForm(
		widget.NewFormItem("Payee", eip.payee),
		widget.NewFormItem("Amount", eip.amount),
		widget.NewFormItem("Memo", eip.memo),
		widget.NewFormItem("Date", eip.date),
	)
	// eip.Form.OnCancel = eip.onCancel
	// eip.Form.OnSubmit = eip.onSubmit
	return eip
}

func (eip *EntryInfoPanel) SetTransaction(t *models.Transaction) {
	eip.transaction = t
	eip.payee.Bind(binding.BindString(&t.Payee))
	amountStr := binding.FloatToStringWithFormat(binding.BindFloat(&t.Amount), "$%5.2f")
	eip.amount.Bind(amountStr)
	eip.memo.Bind(binding.BindString(&t.Memo))
	// eip.date.Bind(binding.BindString(t.Date))
}
