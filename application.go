package main

import (
	//"log"
	//"fmt"
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


var CabinObjects []CabinObject
var Robots []*Robot



func CreateAppInstance(windowDims fyne.Size, banks []*Bank) {

  ApplicationInstance = NewApp(windowDims)

	content := container.NewHBox()

	for _, bank := range banks {
		for _, car := range bank.Cars {

			cabinObj := NewCabinObject(bank.Name, car.Name)
			levels := CabinToLevels(car.Landings)

			cabinObj.elevator = CreateElevatorCabin(windowDims, levels)

			cab := container.NewWithoutLayout(cabinObj.elevator.background)
			cab.Add(cabinObj.elevator.car.container)
			cabinObj.elevator.Place(0)

			if len(CabinObjects) == 0 {
				robot := CreateRobot("Tug", cabinObj.elevator.dimensions.car)
				Robots = append(Robots, robot)
				robot.AssignCar(cabinObj.elevator.car)
				robot.SetFloorState(PCOL_LOBBY)
				cab.Add(robot.image)
				robot.Place(0, cabinObj.elevator.dimensions)
			}

			content.Add(cab)

			CabinObjects = append(CabinObjects, cabinObj)
		}
	}


	windowSize := fyne.NewSize(
		windowDims.Width*float32(len(CabinObjects)),
		windowDims.Height)

	ApplicationInstance.win.SetContent(content)
	ApplicationInstance.win.Resize(windowSize)

	CreateControls(ApplicationInstance.app, banks)

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

	// This is a kludge!
	yOffset = int(windowDims.Height)

	return newApp
}

