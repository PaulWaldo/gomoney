package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

const YYYYMMDD = "2006-01-02"

func MakeTransactionsTable(transactions *[]models.Transaction) *widget.Table {
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

	return table
}

// func (tt transactionTable) updateTransactions() error {
// 	var err error
// 	if ad.selectedAccount == 0 {
// 		ad.Transactions, _, err = ad.Service.Transaction.List()
// 	} else {
// 		ad.Transactions, _, err = ad.Service.Transaction.ListByAccount(ad.selectedAccount)
// 	}
// 	// TODO: call refresh on the table
// 	return err
// }
