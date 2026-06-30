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


//////////////////////////////////////////////////////////////
// Car
//
type Car struct {
	container  *fyne.Container
	objects    *CarObjects

	frontState int
	rearState  int

	yOffset int
	carHeight int
}

func NewCar() Car {
	return Car{
		frontState: DOOR_CLOSED,
		rearState: DOOR_CLOSED,
	}
}

func (c Car) OpenDoor(which int) {
	switch which {

	case FRONT_DOOR:
		if c.frontState == DOOR_OPEN {
			log.Fatal("Cabin door is open when expected to be closed")
		}
		c.objects.front.Hide()
		c.frontState = DOOR_OPEN

	case REAR_DOOR:
		if c.rearState == DOOR_OPEN {
			log.Fatal("Cabin door is open when expected to be closed")
		}
		c.objects.rear.Hide()
		c.rearState = DOOR_OPEN

	default:
		log.Fatal("OpenDoor", which, "is neither front nor rear")
	}
	
}

func (c Car) CloseDoor(which int) {
	switch which {

	case FRONT_DOOR:
		if c.frontState == DOOR_CLOSED {
			log.Fatal("Cabin door is closed when expected to be open")
		}
		c.objects.front.Hide()
		c.frontState = DOOR_CLOSED

	case REAR_DOOR:
		if c.rearState == DOOR_CLOSED {
			log.Fatal("Cabin door is closed when expected to be open")
		}
		c.objects.rear.Hide()
		c.rearState = DOOR_CLOSED

	default:
		log.Fatal("OpenDoor", which, "is neither front nor rear")
	}
}

func (c Car) SetToFloor(floor int) {
	for _, carPos := range CarPositions {
		if carPos.level == floor {
			pos := fyne.NewPos(float32(carPos.xPixCoord), float32(c.yOffset-carPos.yPixCoord-c.carHeight))
			c.container.Move(pos)
			return
		}
	}
}



func CabinCar(floorDims FloorDimensions) *fyne.Container {
	log.Println("Creating a new car")

	car := NewCar()
	car.objects, car.container = CreateCarObjects(floorDims)
	car.carHeight = floorDims.boxHeight
	car.yOffset = yOffset

	return car.container
}
