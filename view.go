package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)


var yOffset int

func DoCanvas(levels []*Level, windowDims fyne.Size) (fyne.App, fyne.Window) {
	log.Println("DoCanvas")

	var graphicsHeight int = int(windowDims.Height)
	var graphicsWidth int = int(windowDims.Width)

	// THIS NEEDS ATTENTION
	yOffset = graphicsHeight
	
	elevatorApp := app.New()
	elevatorWindow := elevatorApp.NewWindow("Cabin")

	elevatorWindow.SetPadded(false)
	elevatorWindow.SetFixedSize(true)

	floorDims := SetDimensions(graphicsHeight, graphicsWidth, len(levels))
	CarPositions = SetCarPositions(levels, floorDims)
	
	vBox := CreateFloors(graphicsHeight, floorDims, levels)
	vBox.Resize(fyne.NewSize(windowDims.Width, windowDims.Height))

	// Create a car instance
	Car := NewCar(floorDims)
	CarInstance.SetToFloor(0)
	
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
}


// Grey background, a rectangle.
//
func CreateBackgroundbox(graphicsHeight float32, graphicsWidth float32) *fyne.Container {

	grey := color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	backgroundbox := canvas.NewRectangle(grey)
	backgroundbox.Resize(fyne.NewSize(graphicsWidth, graphicsHeight))

	return container.NewWithoutLayout(backgroundbox)
}
