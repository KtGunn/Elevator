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
	cab := container.NewWithoutLayout(cabin.background)
	cab.Add(cabin.floors)
	cab.Add(cabin.car)

	cabin2 := CreateElevatorCabin(windowDims, levels)
	cab2 := container.NewWithoutLayout(cabin2.background)
	cab2.Add(cabin2.floors)
	cab2.Add(cabin2.car)

	content := container.NewHBox(
		cab, cab2,
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


