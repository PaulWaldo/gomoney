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

	label := widget.NewLabel("I am a longish label that just sits here")

	buttons := container.NewVBox(
		widget.NewButton("Show", func() {label.Show()}),
		widget.NewButton("Hide", func() {label.Hide()}),
	)

	c := container.NewHSplit(buttons, label)

	myWindow.SetContent(c)
	myWindow.Resize(fyne.NewSize(500, 300))
	myWindow.ShowAndRun()
}
