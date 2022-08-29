package ui

import (
	"fyne.io/fyne/v2"
)

var _ fyne.Layout = (*collapsableLayout)(nil)

type collapsableLayout struct{}

func NewCollapsibleLayout() fyne.Layout {
	return &collapsableLayout{}
}

func (l *collapsableLayout) Layout(items []fyne.CanvasObject, size fyne.Size) {
	if len(items) > 1 {
		panic("more than 1 item")
	}
	i := items[0]
	if size.Width > 0 {
		i.Show()
	}
	if i.Visible() {
		i.Resize(size)
		i.Move(fyne.NewPos(0, 0))
	}
}

func (l *collapsableLayout) MinSize(items []fyne.CanvasObject) fyne.Size {
	if len(items) > 1 {
		panic("more than 1 item")
	}
	i := items[0]
	if i.Visible() {
		return i.MinSize()
	} else {
		return fyne.NewSize(0, 0)
	}
}
