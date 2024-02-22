package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	widget2 "fynebind/widget"
	"image/color"
	"log"
	"time"
)

// doubleLabelUntypedBind creates a widget.DoubleLabel and binds data to it.
// There are two bind data to choose from.
// One of the bind data is continuously updated every 5 seconds.
func doubleLabelUntypedBind() *fyne.Container {
	myData1 := &MyData{
		Name: "Kalle Gurka",
		Info: "Some information",
		Sub:  SubData{Data: "Some data"},
	}

	myData2 := &MyData{
		Name: "Kalle Gurka 2",
		Info: "Some information 2",
		Sub:  SubData{Data: "Some data 2"},
	}

	boundMyData1 := binding.BindUntyped(myData1)
	boundMyData2 := binding.BindUntyped(myData2)

	boundMyData1.AddListener(binding.NewDataListener(func() {
		data, _ := boundMyData1.Get()
		log.Printf("Bound data 1 listen to changes and heard: %+v    (Original struct: %+v)\n", data, myData1)
	}))

	boundMyData2.AddListener(binding.NewDataListener(func() {
		data, _ := boundMyData2.Get()
		log.Printf("Bound data 2 listen to changes and heard: %+v    (Original struct: %+v)\n", data, myData2)
	}))

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			myData1.Sub.Data = "Now is the time " + time.Now().Format("2006-01-02 15:04:05")
			boundMyData1.Reload()
		}
	}()

	doubleLabel := widget2.NewDoubleLabelWithData(boundMyData1, DoubleLabelMyDataContentExtractorFn)

	rebind1LabelButton := widget.NewButton("Rebind label to\nmyData1", func() {
		doubleLabel.Bind(boundMyData1)
	})
	rebind2LabelButton := widget.NewButton("Rebind label to\nmyData2", func() {
		doubleLabel.Bind(boundMyData2)
	})

	captionPanel := createCaptionPanel("Double label widget", "rebinds to different values")
	labelBindContainer := container.NewPadded(container.NewHBox(rebind1LabelButton, rebind2LabelButton, doubleLabel))
	mainPanel := container.NewBorder(captionPanel, nil, nil, nil, labelBindContainer)
	bg := getBackgroundPanel(color.RGBA{R: 16, G: 128, B: 16, A: 24})

	return container.NewStack(bg, mainPanel)
}
