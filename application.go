package main

import (
	//"log"
	"fmt"
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

type CabinObject struct {
	elevator ElevatorCabin
	bank string
	cabin string
}
func NewCabinObject (bank string, cabin string) CabinObject {
	return CabinObject{
		bank: bank,
		cabin: cabin,
	}
}





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
			content.Add(cab)

			CabinObjects = append(CabinObjects, cabinObj)
		}
	}

	robot := CreateRobot("Tug", CabinObjects[0].elevator.dimensions.car)
	Robots = append(Robots, robot)
	robot.AssignCar(CabinObjects[0].elevator.car)

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



func CarFromName(name string) *Car {
	for _, co := range CabinObjects {
		if name == co.cabin {
			return co.elevator.car
		}
	}
	return nil
}

// GetCabinObject
//  returns the CabinObject
func GetCabinObject(bank string, car string) (CabinObject, error) {

	for _, co := range CabinObjects {
		if bank == co.bank {
			return co, nil
		}
		if car == co.cabin {
			return co, nil
		}
	}

	return CabinObject{}, fmt.Errorf("cabin object not found")
}



