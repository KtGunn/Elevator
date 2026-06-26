package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func CreateFloors(ycoordOffset int, floorDims FloorDimensions, levels []*Level) *fyne.Container {
	log.Println("Creating floors.")
	
	vbox := container.NewWithoutLayout()

	for index, level := range levels {

		height := index * floorDims.floorHeight + floorDims.bottomLevel
		flObj := CreateFloorObject(ycoordOffset, height, floorDims, level.Front, level.Rear)

		for _, obj := range flObj {
			vbox.Add(obj)
		}
	}

	return vbox
}

func CreateFloorObject(yOff int, yLevel int, floorDims FloorDimensions, front bool, rear bool) []fyne.CanvasObject {

	objs := []fyne.CanvasObject{}
	position := fyne.Position{X: 0, Y: float32(yLevel)}

	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	grey := color.RGBA{R: 100, G: 100, B: 100, A: 255}

	if front {

		line, pos := CreateALine(black, yOff, position, 2, floorDims.hallLength)
		objs = append(objs, line)
		position = pos

		line, pos = CreateALine(red, yOff, position, 4, floorDims.lobbyLength)
		objs = append(objs, line)
		position = pos

	} else {
		position.X = position.X + float32(floorDims.hallLength+floorDims.lobbyLength)
	}

	line, pos := CreateALine(grey, yOff, position, 2, floorDims.carLength)
	objs = append(objs, line)
	position = pos

	if rear {
		line, pos := CreateALine(red, yOff, position, 4, floorDims.lobbyLength)
		objs = append(objs, line)
		position = pos

		line, pos = CreateALine(black, yOff, position, 2, floorDims.hallLength)
		objs = append(objs, line)
		position = pos
	}

	return objs
}


// Position1 == top-left  Position2 == bottom-right
//
func CreateALine(color color.RGBA, yOff int, position fyne.Position, width int, length int) (*canvas.Line, fyne.Position) {
	anyLine := canvas.NewLine(color)
	anyLine.StrokeWidth =  float32(width)

	anyLine.Position1 = flipVertical(yOff, position)
	position.X = position.X+float32(length)
	anyLine.Position2 = flipVertical(yOff, position)

	return anyLine, position
}

func flipVertical(height int, pos fyne.Position) fyne.Position {
	return fyne.Position{
		X: pos.X,
		Y: float32(height)-pos.Y,
	}
}

