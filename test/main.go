// package main

// import (
// 	"image/color"
// 	"time"

// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// )

// type transaction struct {
// 	Name   string
// 	Amount float64
// 	Date   time.Time
// 	Memo   string
// }

// func main() {
// numTrans := 500
// transactions := make([]transaction, numTrans)
// for i := 0; i < numTrans; i++ {
// 	transactions[i] = transaction{Name: fmt.Sprintf("name %3d", i), Memo: fmt.Sprintf("memo %3d", i), Amount: 123, Date: time.Now()}
// }
// var bindings []binding.DataMap
// for i := range transactions {
// 	_ = append(bindings, binding.BindStruct(&transactions[i]))
// }
// table := widget.NewTable(
// 	func() (int, int) {
// 		return len(transactions), 4
// 	},
// 	func() fyne.CanvasObject {
// 		return widget.NewLabel("wide content")
// 	},
// 	func(i widget.TableCellID, o fyne.CanvasObject) {
// 		switch i.Col {
// 		case 0:
// 			o.(*widget.Label).SetText(transactions[i.Row].Name)
// 		case 1:
// 			o.(*widget.Label).SetText(transactions[i.Row].Date.Format("2006-01-02"))
// 		case 2:
// 			o.(*widget.Label).SetText(fmt.Sprintf("%.2f", transactions[i.Row].Amount))
// 		case 3:
// 			o.(*widget.Label).SetText(transactions[i.Row].Memo)
// 		}
// 	},
// )
// a := app.New()
// w := a.NewWindow("Bug Test")
// w.Resize(fyne.NewSize(600, 400))
// sidebar := container.NewVBox(widget.NewLabel("side"), widget.NewLabel("bar"))
// center := container.NewHSplit(sidebar, table)
// content := container.NewBorder(
// 	nil,
// 	nil,
// 	nil,
// 	nil,
// 	center,
// )
// w.SetContent(content)
// w.ShowAndRun()

// myApp := app.New()
// w := myApp.NewWindow("Rectangle")

// rect := canvas.NewRectangle(color.RGBA{R: 128, A: 128})
// w.SetContent(rect)

// w.Resize(fyne.NewSize(150, 100))
// w.ShowAndRun()
// }

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Layout")

	right := container.NewVBox(
		widget.NewLabelWithStyle("Label1", fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Monospace: true}),
		widget.NewLabel("Label2"),
		widget.NewLabel("Label3"),
	)
	middle := container.NewMax(widget.NewTable(
		func() (int, int) { return 10, 10 },
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {},
	))
	left := widget.NewForm(
		widget.NewFormItem("Payee", widget.NewEntry()),
		widget.NewFormItem("Payee2", widget.NewEntry()),
	)
	rightAndMiddle := container.NewHSplit(right, middle)
	all := container.NewHSplit(rightAndMiddle, left)

	myWindow.SetContent(container.NewMax(all))
	myWindow.ShowAndRun()
}
