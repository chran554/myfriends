package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	widget2 "fynebind/widget"
	"image/color"
	"math/rand"
	"strconv"
	"time"
)

type SubData struct {
	Data string
}

type MyData struct {
	Name string
	Info string
	Sub  SubData
}

func DoubleLabelMyDataContentExtractorFn(boundData binding.ExternalUntyped) (string, string) {
	value, _ := boundData.Get()
	data := value.(MyData)
	return data.Name, data.Sub.Data
}

func main() {
	application := app.New()
	application.SetIcon(theme.ComputerIcon())

	window := application.NewWindow("Untyped binding test")

	doubleLabelUntypedBindContainer := doubleLabelUntypedBind()
	doubleLabelListUntypedBindContainer := doubleLabelListUntypedBind()
	doubleLabelUntypedBindListContainer := doubleLabelUntypedBindInList()

	split := container.NewHSplit(doubleLabelListUntypedBindContainer, doubleLabelUntypedBindListContainer)

	mainLayout := container.NewBorder(nil, doubleLabelUntypedBindContainer, nil, nil, split)

	window.Resize(fyne.NewSize(800, 600))
	window.SetContent(mainLayout)

	window.ShowAndRun()
}

func doubleLabelUntypedBindInList() *fyne.Container {
	var dataList []binding.ExternalUntyped

	for i := 0; i < 5; i++ {
		myData := createMyData(rand.Int())
		dataList = append(dataList, binding.BindUntyped(&myData))
	}

	list := widget.NewList(
		func() int {
			return len(dataList)
		},
		func() fyne.CanvasObject {
			return widget2.NewDoubleLabelWithData(nil, DoubleLabelMyDataContentExtractorFn)
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			boundData := dataList[id]
			object.(*widget2.DoubleLabel).Bind(boundData)
		},
	)

	button1 := widget.NewButton("Change data #1", func() {
		if len(dataList) >= 1 {
			boundValue := dataList[0]
			value, _ := boundValue.Get()
			myData := value.(MyData)
			myData.Sub.Data = time.Now().Format("15:04:05")
			boundValue.Set(myData)
		}
	})

	button2 := widget.NewButton("Change data #2", func() {
		if len(dataList) >= 2 {
			boundValue := dataList[1]
			value, _ := boundValue.Get()
			myData := value.(MyData)
			myData.Sub.Data = time.Now().Format("15:04:05")
			boundValue.Set(myData)
		}
	})

	captionPanel := createCaptionPanel("List with bound items", "items in list listens to changes in bound data")
	buttonPanel := container.NewVBox(button1, button2)
	mainPanel := container.NewPadded(container.NewBorder(captionPanel, buttonPanel, nil, nil, list))
	bg := getBackgroundPanel(color.RGBA{R: 128, G: 16, B: 16, A: 24})

	return container.NewStack(bg, mainPanel)
}

func createCaptionPanel(header, description string) *fyne.Container {
	caption := widget.NewLabel(header)
	caption.TextStyle.Bold = true
	caption.Importance = widget.HighImportance
	subCaption := widget.NewLabel("(" + description + ")")
	subCaption.Importance = widget.HighImportance
	captionPanel := container.NewVBox(caption, subCaption)
	return captionPanel
}

func getBackgroundPanel(rgba color.RGBA) *canvas.Rectangle {
	bg := canvas.NewRectangle(rgba)
	bg.CornerRadius = theme.SelectionRadiusSize()
	bg.StrokeColor = rgba
	bg.StrokeWidth = 2.0
	return bg
}

func createMyData(i int) MyData {
	id := strconv.Itoa(i)
	return MyData{Name: "Name " + id, Info: "Info " + id, Sub: SubData{Data: "Data " + id}}
}
