package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"myfriends/internal/app/resources"
)

type PadLockButton struct {
	widget.BaseWidget
	OnTapped func(locked bool)
	locked   bool
	button   *widget.Button
}

func NewPadLockButton(locked bool, tappedFn func(locked bool)) *PadLockButton {
	padLockButton := &PadLockButton{locked: locked, OnTapped: tappedFn}
	padLockButton.ExtendBaseWidget(padLockButton)
	return padLockButton
}

func (item *PadLockButton) CreateRenderer() fyne.WidgetRenderer {
	item.button = widget.NewButton("", nil)
	item.button.Importance = widget.LowImportance

	item.button.OnTapped = func() {
		item.locked = !item.locked
		item.updateUI()

		if item.OnTapped != nil {
			item.OnTapped(item.locked)
		}
	}

	item.updateUI()

	return widget.NewSimpleRenderer(item.button)
}

func (item *PadLockButton) updateUI() {
	if item.locked {
		item.button.SetIcon(resources.PadLockClosedImageResource)
	} else {
		item.button.SetIcon(resources.PadLockOpenImageResource)
	}
}
