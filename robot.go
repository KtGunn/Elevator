package main

import (
	"log"
	"fyne.io/fyne/v2"
)

const (
	PCOL_DONE    int = 0
	PCOL_RESERVE int = 1
	PCOL_LOBBY   int = 2
	PCOL_ATCAR   int = 3
	PCOL_INCAR   int = 4
	PCOL_OUTCAR  int = 5
)


var Protocol map[int]string = map[int]string{
	PCOL_RESERVE: "Reserve",
	PCOL_LOBBY:   "Lobby",
	PCOL_ATCAR:   "Atcar" ,
	PCOL_INCAR:   "Incar",
	PCOL_OUTCAR:  "Outcar",
	PCOL_DONE:    "Done",
}


func AllStates() []string {
	var states []string
	for _, label := range Protocol {
		states = append(states, label)
	}
	return states
}


func Pcol(state int) string {
	label, ok := Protocol[state]
	if !ok {
		return ""
	}
	return label
}


func ToPcol(state string) int {
	log.Println("ToPcol", state)
	for pcol, label := range Protocol {
		if state == label {
			return pcol
		}
	}
	return -1
}


type RobotState struct {
	car *Car
	floorNow    int
	floorState  int
}

type Robot struct {
	name string
	state *RobotState

	image      *fyne.Container
	objects    *RobotObjects
	dimensions RobotDimensions
}

func NewRobot(name string) *Robot {
	return &Robot{
		name: name,
	}
}

func (r *Robot) OnFloor(floor int) {
	r.state.car = nil
	r.state.floorNow = floor
}


func (r *Robot) WithElevator(elev *Elevator, floor int, pcol int, side int) {
	log.Println("@Place floor=", floor)
	
	if r.state.car == nil {
		Decks.RemoveRobot(r)
		elev.AddRobot(r.image)
	}

	if pcol == PCOL_DONE {
		log.Println(" Removing a robot ")

		elev.RemoveRobot(r.image)
		//Decks.AddRobot(r, floor)

		r.state.car = nil
		r.state.floorNow = floor

		return
	}

	
	dims := elev.dimensions.floor
	x := dims.xPosition(side, pcol)
	y := dims.yPosition(floor) + r.dimensions.bodyHeight
	
	x, y = toCanvasFrame(x,y)
	r.image.Move(fyne.NewPos(float32(x), float32(y)))
}


func CreateRobot(name string, dims CarDimensions) *Robot {
	log.Println("@CreateRobot")

	robot := NewRobot(name)
	robot.state = &RobotState{
		car: nil,
		floorNow: 0,
		floorState: PCOL_DONE,
	}
	robot.objects, robot.image, robot.dimensions = CreateRobotObjects(dims)

	return robot
}
