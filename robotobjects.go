package main

import (
	//"log"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)


const (
	ROBOT_WHEEL    float32 = 0.08
	ROBOT_HEIGHT   float32 = 0.70
	ROBOT_WIDTH    float32 = 0.40
)


type RobotObjects struct {
	body   fyne.CanvasObject
	wheel  fyne.CanvasObject
}

func NewRobotObjects() *RobotObjects {
	return &RobotObjects{
	}
}

func CreateRobotObjects(floorDims FloorDimensions) (*RobotObjects, *fyne.Container) {

	wheelPixels  := ROBOT_WHEEL * float32(floorDims.boxHeight)
	heightPixels := ROBOT_HEIGHT * float32(floorDims.boxHeight)
	widthPixels  := ROBOT_WIDTH * float32(floorDims.carLength)
	
	wheel := canvas.NewCircle(color.RGBA{R: 10, G: 10, B: 10, A: 255})
	wheel.Resize(fyne.NewSize(float32(wheelPixels), float32(wheelPixels)))

	body := canvas.NewRectangle(color.RGBA{R: 210, G: 210, B: 210, A: 255})
	body.Resize(fyne.NewSize(float32(widthPixels), float32(heightPixels)))

	robot := container.NewWithoutLayout()
	robot.Add(body)
	robot.Add(wheel)
	wheel.Move(fyne.NewPos(0.5*float32(wheelPixels), -0.5*float32(wheelPixels)))

	objs := NewRobotObjects()
	objs.body = body
	objs.wheel = wheel

	return objs, robot
}
