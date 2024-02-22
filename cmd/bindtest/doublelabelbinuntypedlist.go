package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	widget2 "fynebind/widget"
	"image/color"
	"math/rand"
	"time"
)

// doubleLabelListUntypedBind creates a list of widget.DoubleLabel.
// The list's backing list is an []interface{} list that can be altered in length and positions.
// The items of the backing list need to be of type interface{} type and not of your specific data type/struct.
// Convert your data into interface{} before inserting into the backing list.
//
// The widget list is then automatically updated on list.Reload() from the backing list.
// Any data changes of the items in the backing list are not guaranteed to be updated in the UI.
func doubleLabelListUntypedBind() *fyne.Container {
	var myDataList []MyData
	for i := 0; i < 2; i++ {
		myDataList = append(myDataList, createMyData(rand.Int()))
	}

	// A bit awkward, the bounded untyped list can only handle interface{} (and set initially to an []interface{} list).
	// There does not seem to be a way to use the raw myDataList (of type []MyData), but it has to be converted into []interface{}.
	interfaceData := make([]interface{}, len(myDataList))
	for i, myData := range myDataList {
		interfaceData[i] = myData
	}

	untypedList := binding.BindUntypedList(&interfaceData)

	// Add a new list entry every 5 seconds to the original (external) list
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for range ticker.C {
			myData := createMyData(rand.Int())
			myData.Sub.Data = time.Now().Format("15:04:05")

			var iMyData interface{}
			iMyData = myData                               // A bit awkward, you need to convert this to an interface{} to be able to add it
			interfaceData = append(interfaceData, iMyData) // Add item to the original (external) list
			untypedList.Reload()                           // Ask the untyped bound list to update
		}
	}()

	list := widget.NewListWithData(
		untypedList,

		func() fyne.CanvasObject {
			return widget2.NewDoubleLabelWithData(nil, DoubleLabelMyDataContentExtractorFn)
		},

		func(item binding.DataItem, object fyne.CanvasObject) {
			untyped := item.(binding.Untyped)   // Cast list item to binding.Untyped
			myDataInterface, _ := untyped.Get() // Get the interface{} value
			myData := myDataInterface.(MyData)  // Cast the interface{} value into the actual MyData value

			untypedMyData := binding.BindUntyped(&myData) // The DoubleLabel widget wants untyped data to be set to it (not that we use that bind functionality later on)

			doubleLabel := object.(*widget2.DoubleLabel)
			doubleLabel.Bind(untypedMyData)
		},
	)

	addButton := widget.NewButton("Add last", func() {
		myData := createMyData(rand.Int())

		var iMyData interface{}
		iMyData = myData                               // A bit awkward, you need to convert this to an interface{} to be able to add it
		interfaceData = append(interfaceData, iMyData) // Add item to the original (external) list
		untypedList.Reload()                           // Ask the untyped bound list to update

		//untypedList.Append(myData)
	})

	removeFirstButton := widget.NewButton("Remove first", func() {
		interfaceData = interfaceData[1:] // Change the size of the original (external) list
		untypedList.Reload()              // Ask the untyped bound list to update
	})

	removeLastButton := widget.NewButton("Remove last", func() {
		interfaceData = interfaceData[:len(interfaceData)-1] // Change the size of the original (external) list
		untypedList.Reload()                                 // Ask the untyped bound list to update
	})

	captionPanel := createCaptionPanel("List with bound data-list", "UI list listens to size changes in data-list")
	buttonPanel := container.NewVBox(addButton, removeFirstButton, removeLastButton)
	mainPanel := container.NewPadded(container.NewBorder(captionPanel, buttonPanel, nil, nil, list))
	bg := getBackgroundPanel(color.RGBA{R: 16, G: 16, B: 128, A: 24})

	return container.NewStack(bg, mainPanel)

}
