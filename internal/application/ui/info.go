package ui

import (
	"fmt"
	"strconv"
	"time"

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
	original    models.Transaction
	Transaction models.Transaction
}

func (eip *EntryInfoPanel) GetSubmitted() models.Transaction {
	return eip.Transaction
}

func (eip *EntryInfoPanel) OnSubmit() {
	// eip.Form.Refresh()
}

func (eip *EntryInfoPanel) OnCancel() {
	fmt.Println("Putting data back")
	eip.SetTransaction(eip.original)
	eip.Form.Refresh()
}

func MakeEntryInfoPanel() *EntryInfoPanel {
	eip := &EntryInfoPanel{
		payee:  widget.NewEntry(),
		amount: widget.NewEntry(),
		memo:   widget.NewMultiLineEntry(),
		date:   widget.NewEntry(),
	}
	// eip.date.SetText(formatDate(eip.date.Text))
	eip.Form = *widget.NewForm(
		widget.NewFormItem("Payee", eip.payee),
		widget.NewFormItem("Amount", eip.amount),
		widget.NewFormItem("Memo", eip.memo),
		widget.NewFormItem("Date", eip.date),
	)
	// eip.date.Validator = dateValidator
	eip.amount.Validator = amountValidator
	return eip
}

// const dateFormat = "02 Jan 06 15:04"
func formatDate(date *time.Time) string {
	if date == nil {
		return ""
	}
	return date.Format(YYYYMMDD)
}

func dateValidator(d string) error {
	_, err := time.Parse(YYYYMMDD, d)
	return err
}

func amountValidator(a string) error {
	_, err := strconv.ParseFloat(a, 64)
	return err
}

func (eip *EntryInfoPanel) SetTransaction(t models.Transaction) {
	eip.original = t
	// Copy the original so that a Cancel will not change the data
	eip.Transaction = eip.original
	eip.payee.Bind(binding.BindString(&eip.Transaction.Payee))
	amountStr := binding.FloatToStringWithFormat(binding.BindFloat(&eip.Transaction.Amount), "%5.2f")
	eip.amount.Bind(amountStr)
	eip.memo.Bind(binding.BindString(&eip.Transaction.Memo))
	// eip.date.Bind(binding.BindString(t.Date))
}
