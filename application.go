package main

import (
	"log"

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


func CreateAppInstance(windowDims fyne.Size, levels []*Level) {

  ApplicationInstance = NewApp(windowDims)
	cabin := CreateElevatorCabin(windowDims, levels)

	content := container.NewVBox(
		cabin,
	)
	ApplicationInstance.win.SetContent(content)
	
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


func DoCanvas(levels []*Level, windowDims fyne.Size) (fyne.App, fyne.Window) {
	log.Println("DoCanvas")

	/*	
	CarContainer = CreateCar(floorDims)
	
	backgroundbox := CreateBackgroundbox(float32(graphicsHeight), float32(graphicsWidth))
	backgroundbox.Add(vBox)
	backgroundbox.Add(Car)
	backgroundbox.Resize(fyne.NewSize(windowDims.Width, windowDims.Height))

	content := container.NewVBox(
		backgroundbox,
	)
	
	elevatorWindow.SetContent(content)
	elevatorWindow.Resize(fyne.NewSize(windowDims.Width, windowDims.Height))

	elevatorApp.Lifecycle().SetOnStarted(func() {
		log.Println("The app has started.")
	})
	
	return elevatorApp, elevatorWindow
	*/
	return nil, nil
}


