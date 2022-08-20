package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

type InfoFormDialog struct {
	Parent      fyne.Window
	Transaction *models.Transaction
}

func (i InfoFormDialog) ShowInfoForm() InfoFormDialog {
	dialog.ShowForm("Update Transaction", "Update", "Cancel",
		[]*widget.FormItem{
			widget.NewFormItem("Payee", widget.NewEntryWithData(binding.BindString(&i.Transaction.Payee))),
		},
		i.a, i.Parent)
	return i
}

func (i InfoFormDialog) a(b bool) {}

type EntryInfoPanel struct {
	Form widget.Form
}

func MakeEntryInfoPanel() *EntryInfoPanel {
	eip := &EntryInfoPanel{}
	payeeEntry := widget.NewEntry()                                                        /*widget.NewEntryWithData(binding.FloatToString(data))*/
	payeeEntry.SetText("MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM") //fmt.Sprintf("%50s", ""))
	eip.Form = *widget.NewForm(
		widget.NewFormItem("Payee", payeeEntry), //widget.NewEntry()/*widget.NewEntryWithData(binding.FloatToString(data))*/),
		widget.NewFormItem("Payee2", widget.NewEntry()),
	)
	return eip
}

// func (eip)
