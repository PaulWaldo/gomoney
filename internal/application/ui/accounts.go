package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/PaulWaldo/gomoney/pkg/domain/models"
)

// type accountList struct {
// 	// accounts *[]models.Account
// 	widget *widget.List
// }

// func (al accountList) onSelected(i widget.ListItemID) {
// 	err := ad.SetSelectedAccount(uint(i))
// 	if err != nil {
// 		fmt.Printf("error setting selected account %d: %s\n", i, err)
// 	}
// }

// func NewAccountList(accounts *[]models.Account) *accountList {
// 	al := &accountList{}
// 	al.makeAccountList(accounts)
// 	return al
// }

func MakeAccountList(accounts *[]models.Account) *widget.List {
	var bindings []binding.DataMap
	for i := range *accounts {
		x := (*accounts)[i]
		bindings = append(bindings, binding.BindStruct(&x))
	}
	list:= widget.NewList(
		func() int {
			return len(*accounts)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("My Checking Account")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			name, err := bindings[i].GetItem("Name")
			if err != nil {
				panic(err)
			}
			x, err := name.(binding.String).Get()
			if err != nil {
				panic(err)
			}
			o.(*widget.Label).SetText(x)
		},
	)
	return list
}

// func (ad AppData) makeLeftSidebar() *fyne.Container {
// 	var bindings []binding.DataMap
// 	for i := range ad.Accounts {
// 		bindings = append(bindings, binding.BindStruct(&ad.Accounts[i]))
// 	}
// 	list := widget.NewList(
// 		func() int {
// 			return len(ad.Accounts)
// 		},
// 		func() fyne.CanvasObject {
// 			return widget.NewLabel("My Checking Account")
// 		},
// 		func(i widget.ListItemID, o fyne.CanvasObject) {
// 			name, err := bindings[i].GetItem("Name")
// 			if err != nil {
// 				panic(err)
// 			}
// 			x, err := name.(binding.String).Get()
// 			if err != nil {
// 				panic(err)
// 			}
// 			o.(*widget.Label).SetText(x)
// 		},
// 	)
// 	list.OnSelected = ad.onSelected
// 	return container.NewMax(list)
// }

// func (ad AppData) SetSelectedAccount(accountId uint) error{
// 	ad.selectedAccount = accountId
// 	return ad.updateTransactions()
// }