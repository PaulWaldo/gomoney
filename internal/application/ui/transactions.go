package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

const YYYYMMDD = "2006-01-02"

type TableSelectedCallback func(i widget.TableCellID)
type TransactionsTable struct {
	Table        *widget.Table
	Selected     *models.Transaction
	Transactions *[]models.Transaction
	// OnSelected   TableSelectedCallback
	mainWindow   fyne.Window
}

func MakeTransactionsTable(transactions *[]models.Transaction, mainWindow fyne.Window) TransactionsTable {
	// https://stackoverflow.com/questions/68085584/binding-table-data-in-go-fyne/73268350#73268350
	// var bindings []binding.DataMap
	// for i := range *transactions {
	// 	bindings = append(bindings, binding.BindStruct(&(*transactions)[i]))
	// }
	table := widget.NewTable(
		func() (int, int) {
			return len(*transactions), 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			switch i.Col {
			// case 0:
			// 	payee, err := bindings[i.Row].GetItem("Payee")
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// 	x, err := payee.(binding.String).Get()
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// 	o.(*widget.Label).SetText(x)
			case 0:
				o.(*widget.Label).SetText((*transactions)[i.Row].Payee)
			case 1:
				o.(*widget.Label).SetText((*transactions)[i.Row].Date.Format(YYYYMMDD))
			case 2:
				o.(*widget.Label).SetText(fmt.Sprintf("%.2f", (*transactions)[i.Row].Amount))
			case 3:
				o.(*widget.Label).SetText((*transactions)[i.Row].Memo)
			}
		},
	)
	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 300)

	tt := TransactionsTable{Table: table, Transactions: transactions, mainWindow: mainWindow}
	// table.OnSelected = tt.OnSelected
	return tt
}

func (tt *TransactionsTable) SetOnSelectedCallback(t TableSelectedCallback) {
	tt.Table.OnSelected = t
}

// func (tt *TransactionsTable) OnSelected(i widget.TableCellID) {
// 	fmt.Printf("Selected %v\n", i)
// 	tt.Selected = &((*tt.Transactions)[i.Row])
// 	d := InfoFormDialog{Parent: tt.mainWindow, Transaction: tt.Selected}
// 	d.ShowInfoForm()
// }
