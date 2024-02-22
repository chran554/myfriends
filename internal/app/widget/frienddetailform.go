package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"myfriends/internal/app/model"
)

type FriendDetailForm struct {
	widget.BaseWidget

	firstNameEntry  *widget.Entry
	familyNameEntry *widget.Entry
	ageLabel        *widget.Label
	ageSlider       *widget.Slider

	boundData                binding.ExternalUntyped
	boundFirstName           binding.String
	boundFamilyName          binding.String
	boundAge                 binding.Int
	boundDataChangedListener binding.DataListener

	editable bool
}

func NewFriendDetailForm() *FriendDetailForm {
	boundFirstName := binding.NewString()
	boundFamilyName := binding.NewString()
	boundAge := binding.NewInt()

	item := &FriendDetailForm{
		firstNameEntry:  widget.NewEntryWithData(boundFirstName),
		familyNameEntry: widget.NewEntryWithData(boundFamilyName),
		ageLabel:        widget.NewLabelWithData(binding.IntToString(boundAge)),
		ageSlider:       widget.NewSliderWithData(0, 130, binding.StringToFloatWithFormat(binding.IntToStringWithFormat(boundAge, "%d.000000"), "%f")),

		boundFirstName:  boundFirstName,
		boundFamilyName: boundFamilyName,
		boundAge:        boundAge,
	}

	item.firstNameEntry.Validator = nil
	item.familyNameEntry.Validator = nil

	// Update the Friend bound data (if form is edited)
	formUpdateListener := binding.NewDataListener(func() {
		value, _ := item.boundData.Get()
		friend := value.(model.Friend)

		friend.FirstName, _ = boundFirstName.Get()
		friend.FamilyName, _ = boundFamilyName.Get()
		friend.Age, _ = boundAge.Get()

		item.boundData.Set(friend)
	})

	// If any form entry is edited, then it will trigger a Friend bound data update
	boundFirstName.AddListener(formUpdateListener)
	boundFamilyName.AddListener(formUpdateListener)
	boundAge.AddListener(formUpdateListener)

	// If Friend bound data is updated externally, then update the form
	boundDataChangedListener := binding.NewDataListener(func() {
		value, _ := item.boundData.Get()
		friend := value.(model.Friend)

		item.boundFirstName.Set(friend.FirstName)
		item.boundFamilyName.Set(friend.FamilyName)
		item.boundAge.Set(friend.Age)

		item.Refresh()
	})

	item.boundDataChangedListener = boundDataChangedListener

	item.ExtendBaseWidget(item)

	return item
}

func (form *FriendDetailForm) Bind(data binding.ExternalUntyped) {
	if form.boundData != nil && form.boundDataChangedListener != nil {
		form.boundData.RemoveListener(form.boundDataChangedListener)
	}

	form.boundData = data

	if form.boundData != nil {
		form.boundData.AddListener(form.boundDataChangedListener)
	}
}

func (form *FriendDetailForm) CreateRenderer() fyne.WidgetRenderer {
	firstNameFormItem := widget.NewFormItem("First name", form.firstNameEntry)
	familyNameFormItem := widget.NewFormItem("Family name", form.familyNameEntry)
	ageFormItem := widget.NewFormItem("Age", form.ageLabel)
	ageChangeFormItem := widget.NewFormItem("", form.ageSlider)

	formWidget := widget.NewForm(firstNameFormItem, familyNameFormItem, ageFormItem, ageChangeFormItem)

	return widget.NewSimpleRenderer(formWidget)
}

func (form *FriendDetailForm) Disabled() bool {
	return !form.editable
}

func (form *FriendDetailForm) Enable() {
	form.SetEditable(true)
}

func (form *FriendDetailForm) Disable() {
	form.SetEditable(false)
}

func (form *FriendDetailForm) SetEditable(editable bool) {
	form.editable = editable
	form.updateEditable()
}

func (form *FriendDetailForm) updateEditable() {
	if form.editable {
		form.firstNameEntry.Enable()
		form.familyNameEntry.Enable()
		form.ageSlider.Show()
	} else {
		form.firstNameEntry.Disable()
		form.familyNameEntry.Disable()
		form.ageSlider.Hide()
	}

	form.Refresh()
}
