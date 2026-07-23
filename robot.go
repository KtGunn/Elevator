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
	PCOL_LOBBY: "Lobby",
	PCOL_ATCAR: "Atcar" ,
	PCOL_INCAR: "Incar",
	PCOL_OUTCAR: "Outcar",
	PCOL_DONE: "Done",
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

func (r *Robot) AssignCar(car *Car) {
	r.state.car = car
}

func (r *Robot) SetFloorState(state int) {
	log.Println("@SetFloorState state=", state)
	r.state.floorState = state
}

func (r *Robot) Place(floor int, dims ElevatorDimensions) {
	log.Println("@Place floor=", floor)
	
	xPix, yPix := r.positionAt(floor, dims)
	r.image.Move(fyne.NewPos(xPix, float32(yOffset)-yPix))
}

func (r *Robot) positionAt(floor int, dims ElevatorDimensions) (float32, float32) {
	log.Println("@positionAt floor=", floor)

	floorY := floor*dims.floor.floorHeight + dims.floor.bottomLevel
	bodyH := float32(r.dimensions.bodyHeight)
	bodyW := float32(r.dimensions.bodyWidth)

	var x float32
	switch r.state.floorState {
	case PCOL_INCAR:
		log.Println("positionAt INCAR")
		x = float32(dims.floor.hallLength + dims.floor.lobbyLength)
	case PCOL_ATCAR:
		log.Println("positionAt ATCAR")
		x = float32(dims.floor.hallLength+dims.floor.lobbyLength) - bodyW
	default:
		log.Println("positionAt default", r.state.floorState)
		x = float32(dims.floor.hallLength) + (float32(dims.floor.lobbyLength)-bodyW)/2
	}

	return x, float32(floorY) + bodyH
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
