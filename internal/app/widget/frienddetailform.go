package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fynebind/internal/app/model"
	widget2 "fynebind/pkg/widget"
	"strconv"
)

type FriendListItem struct {
	widget.BaseWidget

	firstName  *widget.Label
	familyName *widget.Label
	age        *widget.Label

	boundData                binding.ExternalUntyped
	boundDataChangedListener binding.DataListener
	bindTextFn               widget2.DoubleLabelBindTextFn
}

func NewFriendListItem(boundData binding.ExternalUntyped) *FriendListItem {
	item := &FriendListItem{
		firstName:  widget.NewLabel(""),
		familyName: widget.NewLabel(""),
		age:        widget.NewLabel(""),
	}

	boundDataChangedListener := binding.NewDataListener(func() {
		value, _ := item.boundData.Get()
		friend := value.(model.Friend)

		item.firstName.Text = friend.FirstName
		item.familyName.Text = friend.FamilyName
		item.age.Text = "(" + strconv.Itoa(friend.Age) + ")"

		item.Refresh()
	})
	item.boundDataChangedListener = boundDataChangedListener

	item.Bind(boundData)

	item.ExtendBaseWidget(item)

	return item
}

func (item *FriendListItem) Bind(data binding.ExternalUntyped) {
	if item.boundData != nil && item.boundDataChangedListener != nil {
		item.boundData.RemoveListener(item.boundDataChangedListener)
	}

	item.boundData = data

	if item.boundData != nil {
		item.boundData.AddListener(item.boundDataChangedListener)
	}
}

func (item *FriendListItem) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewHBox(item.firstName, item.familyName, layout.NewSpacer(), item.age)
	return widget.NewSimpleRenderer(c)
}
