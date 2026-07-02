package main

import (
	//"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)


var yOffset int

var ApplicationInstance Application

type Application struct {
	app  fyne.App
	win  fyne.Window
	dims fyne.Size
}

func NewApplication() Application {
	return Application{}
}


var Cabins []ElevatorCabin

func CreateAppInstance(windowDims fyne.Size, levels []*Level) {

  ApplicationInstance = NewApp(windowDims)

	cabin := CreateElevatorCabin(windowDims, levels)
	Cabins = append(Cabins, cabin)
	
	cab := container.NewWithoutLayout(cabin.background)
	cab.Add(cabin.car.container)
	cabin.Place(1, windowDims.Height)

	content := container.NewHBox(
		cab,
	)

	ApplicationInstance.win.SetContent(content)
	ApplicationInstance.win.Resize(fyne.NewSize(windowDims.Width*2, windowDims.Height))
	ApplicationInstance.win.ShowAndRun()
}

func NewApp(windowDims fyne.Size) Application {

	newApp := NewApplication()
	newApp.dims = windowDims
	
	
	newApp.app= app.New()
	newApp.win = newApp.app.NewWindow("Cabin")
	newApp.win.Resize(fyne.NewSize(windowDims.Width, windowDims.Height))
	
	newApp.win.SetPadded(false)
	newApp.win.SetFixedSize(true)
	
	return newApp
}




