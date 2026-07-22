package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"

	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateControls(app fyne.App, banks []*Bank) {

	win := app.NewWindow("Controls")

	width := 200
	height := 600
	win.Resize(fyne.NewSize(float32(width), float32(height)))

	cabinSide, cabinSelector, floorSelector := CabinControls(app, banks)
	robotSide := RobotControls(app, banks, cabinSelector, floorSelector)

	win.SetContent(container.NewHBox(
		cabinSide,
		robotSide,
	))

	win.Resize(fyne.NewSize(200, 200))
	win.Show()

}

func CabinControls(app fyne.App, banks []*Bank) (*fyne.Container, *widget.Select, *widget.Select) {

	var cabinSelector *widget.Select

	floorSelector := widget.NewSelect(nil, func(picked string) {
		fmt.Println("selected floor:", picked)

		cobj, err := GetCabinObject("", cabinSelector.Selected)
		if err != nil {
			fmt.Println("error getting cabin object:", err)
			return
		}

		floor, err := strconv.Atoi(picked)
		if err != nil {
			fmt.Println("error converting floor to int:", err)
			return
		}

		fmt.Println("moving cabin", cobj.cabin, "to floor", floor)
		cobj.elevator.Place(floor)
	})
	floorSelector.PlaceHolder = "Pick a floor"

	cabinSelector = widget.NewSelect(CabinNames(banks), func(picked string) {
		// Set floor selectors
		floors := CabinFloors(picked, banks)
		floorSelector.Options = floors
		floorSelector.ClearSelected()
		floorSelector.PlaceHolder = "Move to floor"
		floorSelector.Refresh()
	})
	cabinSelector.PlaceHolder = "Pick a cabin"

	var ops []string = []string{
		"Open Front Door",
		"Close Front Door",
		"Open Rear Door",
		"Close Rear Door",
	}
	opSelector := widget.NewSelect(ops, func(picked string) {
		fmt.Println("Door action ..")

		cobj, err := GetCabinObject("", cabinSelector.Selected)
		if err != nil {
			fmt.Println("error getting cabin object:", err)
			return
		}
		// VERIFY
		fmt.Println("  bank", cobj.bank, " cabin", cobj.cabin)

		door, action := DoorAndAction(picked)

		switch action {
		case DOOR_OPEN:
			cobj.elevator.car.OpenDoor(door)
		case DOOR_CLOSED:
			cobj.elevator.car.CloseDoor(door)
		}

	})

	opSelector.PlaceHolder = "Door op"

	return container.NewVBox(
		cabinSelector,
		opSelector,
		floorSelector,
	), cabinSelector, floorSelector

}

func RobotControls(app fyne.App, banks []*Bank,
	cabinSelector *widget.Select,
	floorSelector *widget.Select) *fyne.Container {

	var robotSelector *widget.Select
	var stateSelector *widget.Select

	// PCOL state
	//
	stateSelector = widget.NewSelect(AllStates(), func(picked string) {

		if picked == "" {
			return // Prevent infinite loop when ClearSelected is called
		}

		robotName := robotSelector.Selected
		var robot *Robot
		if robot = RobotFromName(robotName); robot == nil {
			fmt.Println("@stateSelector:", robotName, "is invalid. Bye...")
			stateSelector.ClearSelected()
			return
		}

		robot.SetFloorState(ToPcol(picked))
		fmt.Println("selected state:", picked, "or", ToPcol(picked))
	})
	stateSelector.PlaceHolder = "?"

	// ROBOT
	//
	robotSelector = widget.NewSelect(RobotNames(Robots), func(robotName string) {
		fmt.Println("Selected Cabin from CabinControls:", cabinSelector.Selected)
		fmt.Println("Selected Floor from CabinControls:", floorSelector.Selected)

		// robot --> car --> floor --> state
		robot := RobotFromName(robotName)

		floor, err := strconv.Atoi(floorSelector.Selected)
		if err != nil {
			fmt.Println("floor conversion error:", err, "Setting to 0")
			floor = 0
		}

		pcolInt := ToPcol(stateSelector.Selected)
		if pcolInt == -1 {
			fmt.Println("Invalid state selected:", stateSelector.Selected)
			return
		}

		cobj, err := GetCabinObject("", cabinSelector.Selected)
		if err != nil {
			fmt.Println("error getting cabin object:", err, "Bye...")
			return
		}

		fmt.Println(robotName, floor, pcolInt)
		robot.Place(floor, cobj.elevator.dimensions)

	})
	robotSelector.PlaceHolder = "Pick robot"

	return container.NewVBox(
		robotSelector,
		stateSelector,
	)
}

func DoorAndAction(label string) (door int, op int) {

	if strings.Contains(label, "Front") {
		door = FRONT_DOOR
	} else {
		door = REAR_DOOR
	}

	if strings.Contains(label, "Open") {
		op = DOOR_OPEN
	} else {
		op = DOOR_CLOSED
	}

	return door, op
}

func RobotNames(robots []*Robot) []string {
	var names []string
	for _, robot := range robots {
		names = append(names, robot.name)
	}
	return names
}

func CabinNames(banks []*Bank) []string {
	var names []string
	for _, b := range banks {
		for _, c := range b.Cars {
			names = append(names, c.Name)
		}
	}
	return names
}

func CabinFloors(cabin string, banks []*Bank) []string {
	var floors []string

	for _, b := range banks {
		for _, c := range b.Cars {
			if c.Name == cabin {
				for _, l := range c.Landings {
					floors = append(floors, fmt.Sprintf("%d", l.Floor))
				}
			}
		}
	}

	return floors
}
