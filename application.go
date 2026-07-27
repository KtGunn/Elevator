package main

import (
	//"log"
	"fmt"
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
var Decks *Floors
var yOffset int

func toCanvasFrame(x int, y int) (int, int) {
	return x, yOffset - y
}

func AddFloors(lay *fyne.Container, win fyne.Size) *Floors{
	floors := NewFloors()
	number := NumberOfFloors()

	floors.Dimensions(number, int(win.Width), int(win.Height))
	floors.Image()
	lay.Add(floors.image)

	return floors
}

func CreateAppInstance(windowDims fyne.Size, banks []*Bank) {

  ApplicationInstance = NewApp(windowDims)
	content := container.NewHBox()

	fsize := fyne.Size{
		Width: windowDims.Width/2,
		Height: windowDims.Height,
	}
	Decks = AddFloors(content, fsize)

	for _, bank := range banks {
		for _, car := range bank.Cars {

			elev := AddElevator(bank.Name, car.Name, car.Landings, windowDims)
			Elevators = append(Elevators, elev)

			cont := container.NewWithoutLayout()
			cont.Add(elev.image)
			cont.Add(elev.car.image)
			content.Add(cont)
		}
	}

	AddRobots(content)

	windowSize := fyne.NewSize(
		windowDims.Width*(float32(len(Elevators)) + 0.5),
		windowDims.Height,
	)

	ApplicationInstance.win.SetContent(content)
	ApplicationInstance.win.Resize(windowSize)

	CreateControls(ApplicationInstance.app, banks)

	ApplicationInstance.win.ShowAndRun()
}

func AddRobots(appLayout *fyne.Container) {

	for n, elev := range Elevators {

		robot := CreateRobot(fmt.Sprintf("Tug-%d", n), elev.dimensions.car)
		Robots = append(Robots, robot)

		Decks.AddRobot(robot, 0)
	}
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

	return elevator
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

func GetElevator(name string) *Elevator {
	for _,el := range Elevators {
		if el.car.name == name {
			return el
		}
	}
	return nil
}

func RobotFromName(name string) *Robot {
	for _, r := range Robots {
		if name == r.name {
			return r
		}
	}
	return nil
}

