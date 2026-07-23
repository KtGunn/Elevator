package main

import (
	//"log"
	//"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)


type Application struct {
	app  fyne.App
	win  fyne.Window
	dims fyne.Size
}

var ApplicationInstance Application

func NewApplication() Application {
	return Application{}
}


var Elevators []*Elevator
var Robots []*Robot
var yOffset int

func CreateAppInstance(windowDims fyne.Size, banks []*Bank) {

  ApplicationInstance = NewApp(windowDims)
	content := container.NewHBox()

	for _, bank := range banks {
		for _, car := range bank.Cars {

			cabinObj := AddElevator(bank.Name, car.Name, car.Landings, windowDims)
			Elevators = append(Elevators, cabinObj)

			cont := container.NewWithoutLayout()
			cont.Add(cabinObj.image)
			cont.Add(cabinObj.car.image)
			content.Add(cont)
		}
	}

	// AddRobots()

	windowSize := fyne.NewSize(
		windowDims.Width*float32(len(Elevators)),
		windowDims.Height,
	)

	ApplicationInstance.win.SetContent(content)
	ApplicationInstance.win.Resize(windowSize)

	//CreateControls(ApplicationInstance.app, banks)

	ApplicationInstance.win.ShowAndRun()
}

func AddElevator(bank string, car string, landings []*Landing,
	dims fyne.Size) *Elevator {

	floors := NumberOfFloors()
	elevator := NewElevator(bank)

	elevator.Dimension(dims, floors)
	elevator.Levels(landings)
	elevator.Image(dims)

	elevator.Car(car)
	elevator.SetCar(0)

	/*
	elevator := NewElevator(bank, car)
	levels := CabinToLevels(landings)

	elevator.elevator = CreateElevatorCabin(dims, levels)

	image := container.NewWithoutLayout(elevator.elevator.background)
	image.Add(elevator.elevator.car.container)
	elevator.elevator.Place(0)

	elevator.image = image
	*/
	
	return elevator
}

func AddRobots() {

	/*
	for n, cabinObj := range Elevators {

		robot := CreateRobot(fmt.Sprintf("Tug-%d", n), cabinObj.elevator.dimensions.car)
		Robots = append(Robots, robot)

		robot.AssignCar(cabinObj.elevator.car)
		robot.SetFloorState(PCOL_LOBBY)

		cabinObj.image.Add(robot.image)
		robot.Place(0, cabinObj.elevator.dimensions)
	}
	*/
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

func RobotFromName(name string) *Robot {
	for _, r := range Robots {
		if name == r.name {
			return r
		}
	}
	return nil
}

func CabinObjFromCar(car *Car) Elevator {
	/*
	for _, cobj := range Elevators {
		if cobj.elevator.car == car {
			return cobj
		}
	}
	*/
	return Elevator{}
}
