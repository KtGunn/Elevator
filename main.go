package main

import (
	"log"
	"fyne.io/fyne/v2"
)


var ApplicationInstance Application

type Application struct {
	app fyne.App
	win fyne.Window
	dims fyne.Size
}

var Levels []*Level


func main() {
	var err error
	Levels, err = ReadLevels("levels.json")
	if err != nil {
		log.Fatal("could not read json file: ", err)
	}

	ShowLevels(Levels)
	
	ApplicationInstance.dims = fyne.Size{
		Width: 240,
		Height: 280,}

	ApplicationInstance.app, ApplicationInstance.win =
		DoCanvas(Levels, ApplicationInstance.dims)
	
	ApplicationInstance.win.ShowAndRun()

	return
}



