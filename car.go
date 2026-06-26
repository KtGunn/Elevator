package main

import (
	"log"
	"fyne.io/fyne/v2"
)

const (
	FRONT_DOOR int = 0
	FREAR  int = 1
)

const (
	DOOR_OPEN int = 0
	DOOR_CLOSED int = 1
)

type Car struct {
	container  *fyne.Container

	frontState int
	rearState  bool

	yOffset int
	carHeight int
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

var CarInstance Car

func NewCar(floorDims FloorDimensions) *fyne.Container {
	log.Println("Creating a new car")

	car := Car{}
	car.container = CreateCar(floorDims)
	car.carHeight = floorDims.boxHeight
	car.yOffset = yOffset
	CarInstance = car

	return car.container
}
