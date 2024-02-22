package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
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
