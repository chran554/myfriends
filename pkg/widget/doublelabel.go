package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// DoubleLabelBindTextFn takes a generic bound data and extracts/returns the two texts to be displayed in the DoubleLabel.
type DoubleLabelBindTextFn func(boundData binding.ExternalUntyped) (label1Text string, label2Text string)

type DoubleLabel struct {
	widget.BaseWidget

	text1 *widget.Label
	text2 *widget.Label

	boundData                binding.ExternalUntyped
	boundDataChangedListener binding.DataListener
	bindTextFn               DoubleLabelBindTextFn
}

func NewDoubleLabelWithData(boundData binding.ExternalUntyped, bindTextFn DoubleLabelBindTextFn) *DoubleLabel {
	item := &DoubleLabel{
		text1:      widget.NewLabel(""),
		text2:      widget.NewLabel(""),
		bindTextFn: bindTextFn,
	}

	boundDataChangedListener := binding.NewDataListener(func() {
		item.text1.Text, item.text2.Text = item.bindTextFn(item.boundData)
		item.Refresh()
	})
	item.boundDataChangedListener = boundDataChangedListener

	item.Bind(boundData)

	item.ExtendBaseWidget(item)

	return item
}

func (item *DoubleLabel) Bind(data binding.ExternalUntyped) {
	if item.boundData != nil && item.boundDataChangedListener != nil {
		item.boundData.RemoveListener(item.boundDataChangedListener)
	}

	item.boundData = data

	if item.boundData != nil {
		item.boundData.AddListener(item.boundDataChangedListener)
	}
}

func (item *DoubleLabel) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(item.text1, item.text2)
	return widget.NewSimpleRenderer(c)
}
