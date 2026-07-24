package main

import (
	"log"
	"fyne.io/fyne/v2"
)

const (
	FRONT_DOOR int = 0
	REAR_DOOR  int = 1

	DOOR_OPEN int = 0
	DOOR_CLOSED int = 1
)


type Car struct {

	name string
	state *CarState

	image *fyne.Container
	objects *CarObjects
	dimensions CarDimensions
}

func NewCar(name string) *Car {
	return &Car{
		name: name,
	}
}

func (c *Car) OpenDoor(which int) {
	switch which {

	case FRONT_DOOR:
		c.objects.front.Hide()
		c.state.frontOpen = DOOR_OPEN

	case REAR_DOOR:
		c.objects.rear.Hide()
		c.state.rearOpen = DOOR_OPEN

	default:
		log.Fatal("OpenDoor", which, "is neither front nor rear")
	}
}

func (c *Car) CloseDoor(which int) {
	switch which {

	case FRONT_DOOR:
		c.objects.front.Show()
		c.state.frontOpen = DOOR_CLOSED

	case REAR_DOOR:
		c.objects.rear.Show()
		c.state.rearOpen = DOOR_CLOSED

	default:
		log.Fatal("OpenDoor", which, "is neither front nor rear")
	}
}


func (c *Car) SetToFloor(floor int) {
	return
}

type CarState struct {
	floor int
	available bool

	frontOpen int
	rearOpen  int
}


func CreateCar(name string, dims CarDimensions) *Car {
	
	car := NewCar(name)
	car.state = &CarState{}

	car.objects, car.image = CreateCarObjects(dims)
	return car
}
