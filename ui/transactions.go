package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)



func (ad AppData) makeCenter() *fyne.Container {
	var bindings []binding.DataMap
	for i := range ad.Transactions {
		bindings = append(bindings, binding.BindStruct(&ad.Transactions[i]))
	}
	table := widget.NewTable(
		func() (int, int) {
			return len(ad.Transactions), 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			switch i.Col {
			case 0:
				payee, err := bindings[i.Row].GetItem("Payee")
				if err != nil {
					panic(err)
				}
				x, err := payee.(binding.String).Get()
				if err != nil {
					panic(err)
				}
				o.(*widget.Label).SetText(x)
			case 1:
				o.(*widget.Label).SetText(ad.Transactions[i.Row].Date.Format(YYYYMMDD))
			case 2:
				o.(*widget.Label).SetText(fmt.Sprintf("%.2f", ad.Transactions[i.Row].Amount))
			case 3:
				o.(*widget.Label).SetText(ad.Transactions[i.Row].Memo)
			}
		},
	)
	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 300)
	split := container.NewHSplit(ad.makeLeftSidebar(), table)
	split.Offset = 0.2

	return container.NewMax(split)
}

func (ad *AppData) updateTransactions() error {
	var err error
	if ad.selectedAccount == 0 {
		ad.Transactions, _, err = ad.Service.Transaction.List()
	} else {
		ad.Transactions, _, err = ad.Service.Transaction.ListByAccount(ad.selectedAccount)
	}
	// TODO: call refresh on the table
	return err
}
