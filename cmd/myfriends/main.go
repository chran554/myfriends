package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	app2 "myfriends/internal/app"
)

func main() {
	application := app.New()
	application.SetIcon(theme.ComputerIcon())

	window := application.NewWindow("My Friends - an untyped bind test application")
	window.Resize(fyne.NewSize(800, 600))

	window.SetContent(app2.ApplicationContent())

	window.ShowAndRun()
}
