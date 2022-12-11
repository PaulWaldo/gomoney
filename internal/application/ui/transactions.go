package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/fyne-headertable/headertable"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

const YYYYMMDD = "2006-01-02"

type TableSelectedCallback func(i widget.TableCellID)
type TransactionsTable struct {
	/*Table*/ *headertable.SortingHeaderTable
	Selected *models.Transaction
	// Transactions *[]models.Transaction
	// OnSelected   TableSelectedCallback
	mainWindow fyne.Window
	Bindings   []binding.Struct
}

func MakeTransactionsTable(transactions *[]models.Transaction, mainWindow fyne.Window) TransactionsTable {
	bindings := make([]binding.Struct, len(*transactions))

	for i := 0; i < len(*transactions); i++ {
		bindings[i] = binding.BindStruct(&(*transactions)[i])
	}
	to := headertable.TableOpts{
		ColAttrs: []headertable.ColAttr{
			{
				Name:   "Payee",
				Header: "Payee",
				HeaderStyle: headertable.CellStyle{
					Alignment: fyne.TextAlignCenter,
					TextStyle: fyne.TextStyle{Bold: true},
				},
				WidthPercent: 100,
			},
			{
				Name:   "Type",
				Header: "Type",
				HeaderStyle: headertable.CellStyle{
					Alignment: fyne.TextAlignCenter,
					TextStyle: fyne.TextStyle{Bold: true},
				},
				WidthPercent: 50,
			},
			{
				Name:   "Amount",
				Header: "Amount",
				HeaderStyle: headertable.CellStyle{
					Alignment: fyne.TextAlignCenter,
					TextStyle: fyne.TextStyle{Bold: true},
				},
				DataStyle: headertable.CellStyle{
					Alignment: fyne.TextAlignTrailing,
				},
				WidthPercent: 65,
				Converter:    headertable.DisplayAsCurrency,
			},
			{
				Name:   "Memo",
				Header: "Memo",
				HeaderStyle: headertable.CellStyle{
					Alignment: fyne.TextAlignCenter,
					TextStyle: fyne.TextStyle{Bold: true},
				},
				WidthPercent: 120,
			},
			{
				Name:   "Date",
				Header: "Date",
				HeaderStyle: headertable.CellStyle{
					Alignment: fyne.TextAlignCenter,
					TextStyle: fyne.TextStyle{Bold: true},
				},
				DataStyle: headertable.CellStyle{
					Alignment: fyne.TextAlignCenter,
				},
				WidthPercent: 65,
				Converter:    headertable.DisplayAsISODate,
			},
			{
				Name:   "Balance",
				Header: "Balance",
				HeaderStyle: headertable.CellStyle{
					Alignment: fyne.TextAlignCenter,
					TextStyle: fyne.TextStyle{Bold: true},
				},
				DataStyle: headertable.CellStyle{
					Alignment: fyne.TextAlignTrailing,
				},
				WidthPercent: 60,
				Converter:    headertable.DisplayAsCurrency,
			},
		},
		Bindings: bindings,
		RefWidth: "XXXXXXXXXXXXXXXXXXXXXXX",
	}
	st := headertable.NewSortingHeaderTable(&to)

	tt := TransactionsTable{SortingHeaderTable: st, Bindings: bindings, mainWindow: mainWindow}
	return tt
}

func (tt *TransactionsTable) UpdateTransactions(transactions *[]models.Transaction) {
	bindings := make([]binding.Struct, len(*transactions))

	for i := 0; i < len(*transactions); i++ {
		bindings[i] = binding.BindStruct(&(*transactions)[i])
	}
	tt.SortingHeaderTable.TableOpts.Bindings = bindings
	tt.SortingHeaderTable.Refresh()
}

func (tt *TransactionsTable) SetOnSelectedCallback(t TableSelectedCallback) {
	tt.Data.OnSelected = t
}
