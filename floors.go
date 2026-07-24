package main

import (
	//"log"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
)



type Floors struct {

	occupation map[int][]*Robot

	floors      int
	floorHeight int
	bottomLevel int

	width  int
	height int
	
	image *fyne.Container
}

func NewFloors() *Floors {
	return &Floors{
		occupation: make(map[int][]*Robot),
	}
}

func (f *Floors) Dimensions(floors int, width int, height int) {
	
	f.width = width
	f.height = height
	f.floors = floors

	f.floorHeight, _, f.bottomLevel = FloorAndCabHeights(
		f.height,
		f.width,
		f.floors,
	)
}


func (f *Floors) Image() {


	f.image = container.NewWithoutLayout()

	// Background rectangle
	size := fyne.Size{
		Width: float32(f.width),
		Height: float32(f.height),
	}

	imagebox := canvas.NewRectangle(DARK)
	imagebox.Resize(size)
	imagebox.SetMinSize(size)

	f.image.Add(imagebox)

	for pi := range f.floors {

		anyLine := canvas.NewLine(GREY)
		anyLine.StrokeWidth = float32(1)

		x := 0
		y := f.bottomLevel + f.floorHeight * pi
		xp , yp := toCanvasFrame(x, y)
		
		anyLine.Position1 = fyne.NewPos(float32(xp), float32(yp))
		anyLine.Position2 = fyne.NewPos(float32(xp + f.width), float32(yp))

		f.image.Add(anyLine)

		x = 4
		y = (3*f.floorHeight)/4 + f.bottomLevel + f.floorHeight * pi
		xp , yp = toCanvasFrame(x, y)

		label := canvas.NewText(fmt.Sprintf("floor %d", pi), BLACK)
		label.Move(fyne.NewPos(float32(xp), float32(yp)))

		f.image.Add(label)
	}
}
