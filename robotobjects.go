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


func NewRobotDimensions() RobotDimensions {
	return RobotDimensions{}
}

type RobotObjects struct {
	body   fyne.CanvasObject
	wheel  fyne.CanvasObject
}

func NewRobotObjects() *RobotObjects {
	return &RobotObjects{
	}
}

func CreateRobotObjects(cardims CarDimensions) (*RobotObjects, *fyne.Container, RobotDimensions) {

	wheelPixels  := ROBOT_WHEEL * float32(cardims.boxHeight)
	heightPixels := ROBOT_HEIGHT * float32(cardims.boxHeight)
	widthPixels  := ROBOT_WIDTH * float32(cardims.carLength)
	
	wheel := canvas.NewCircle(color.RGBA{R: 10, G: 10, B: 10, A: 255})
	wheel.Resize(fyne.NewSize(float32(wheelPixels), float32(wheelPixels)))

	body := canvas.NewRectangle(color.RGBA{R: 10, G: 110, B: 210, A: 255})
	body.Resize(fyne.NewSize(float32(widthPixels), float32(heightPixels)))

	robot := container.NewWithoutLayout()
	robot.Add(body)
	robot.Add(wheel)
	wheel.Move(fyne.NewPos(0.5*float32(widthPixels-wheelPixels), float32(heightPixels)))
	//wheel.Move(fyne.NewPos(0.5*float32(widthPixels-wheelPixels), -0.5*float32(wheelPixels)))

	objs := NewRobotObjects()
	objs.body = body
	objs.wheel = wheel

	dims := RobotDimensions{
		bodyHeight: int(heightPixels),
		bodyWidth: int(widthPixels),
		wheelDia: int(wheelPixels),
	}

	return objs, robot, dims
}
