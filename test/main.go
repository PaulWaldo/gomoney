package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type transaction struct {
	Name   string
	Amount float64
	Date   time.Time
	Memo   string
}

func main() {
	numTrans := 500
	transactions := make([]transaction, numTrans)
	for i := 0; i < numTrans; i++ {
		transactions[i] = transaction{Name: fmt.Sprintf("name %3d", i), Memo: fmt.Sprintf("memo %3d", i), Amount: 123, Date: time.Now()}
	}
	var bindings []binding.DataMap
	for i := range transactions {
		_ = append(bindings, binding.BindStruct(&transactions[i]))
	}
	table := widget.NewTable(
		func() (int, int) {
			return len(transactions), 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			switch i.Col {
			case 0:
				o.(*widget.Label).SetText(transactions[i.Row].Name)
			case 1:
				o.(*widget.Label).SetText(transactions[i.Row].Date.Format("2006-01-02"))
			case 2:
				o.(*widget.Label).SetText(fmt.Sprintf("%.2f", transactions[i.Row].Amount))
			case 3:
				o.(*widget.Label).SetText(transactions[i.Row].Memo)
			}
		},
	)
	a := app.New()
	w := a.NewWindow("Bug Test")
	w.Resize(fyne.NewSize(600, 400))
	sidebar := container.NewVBox(widget.NewLabel("side"), widget.NewLabel("bar"))
	center := container.NewHSplit(sidebar, table)
	content := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		center,
	)
	w.SetContent(content)
	w.ShowAndRun()
}
