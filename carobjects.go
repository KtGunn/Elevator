package main

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)



type CarObjects struct {
	box   fyne.CanvasObject
	front fyne.CanvasObject
	rear  fyne.CanvasObject
}

func NewCarObjects() *CarObjects {
	return &CarObjects{}
}


func CreateCarObjects(cardims CarDimensions) (*CarObjects, *fyne.Container) {
	box := container.NewWithoutLayout()
	var doorWidth float32 = 8

	car := NewCarObjects()
	grey := color.RGBA{R: 110, G: 110, B: 110, A: 255}
	car.box = canvas.NewRectangle(grey)
	car.box.Resize(fyne.NewSize(float32(cardims.carLength), float32(cardims.boxHeight)))

	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	frontLine := canvas.NewLine(black)
	frontLine.Position1 = fyne.NewPos(doorWidth/2, 8)
	frontLine.Position2 = fyne.NewPos(doorWidth/2, float32(cardims.boxHeight))
	frontLine.StrokeWidth = doorWidth
	car.front = frontLine

	rearLine := canvas.NewLine(black)

	rearLine.Position1 = fyne.NewPos(float32(cardims.carLength)-doorWidth/2, 8)
	rearLine.Position2 = fyne.NewPos(float32(cardims.carLength)-doorWidth/2, float32(cardims.boxHeight))
	rearLine.StrokeWidth = doorWidth
	car.rear = rearLine

	box.Add(car.box)
	box.Add(car.front)
	box.Add(car.rear)

	return car, box
}
