package main

import (
	"log"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var (
	RED   = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	BLACK = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	GREY  = color.RGBA{R: 100, G: 100, B: 100, A: 255}
	DARK  = color.RGBA{R: 200, G: 200, B: 200, A: 255}
)


//////////////////////////////////////////////////////////////
// ElevatorCabin
//
type ElevatorCabin struct {
	car Car
	background *fyne.Container
	dimensions ElevatorDimensions
}

func NewElevatorCabin() ElevatorCabin {
	return ElevatorCabin{}
}



//////////////////////////////////////////////////////////////
// CarPositions
// holds xy-pixel coords of car position at floors
var CarPositions []CarPosition


//////////////////////////////////////////////////////////////
// CreateElevatorCabin
// creates all objects needed to render an elevator cabin
//
//  NewDims() returns {floor: FloorDimensions, positions: []CarPosition}

func CreateElevatorCabin(dims fyne.Size, levels []*Level) ElevatorCabin {
	log.Println("New elevator cabin")

	newCabin := NewElevatorCabin()
	newCabin.dimensions = ElevatorDims(dims,levels)

	newCabin.background = Background(dims, newCabin.dimensions.floor, levels)
	newCabin.car = CabinCar(newCabin.dimensions.floor)

	return newCabin
}

func ElevatorDims(winDims fyne.Size, floors []*Level) ElevatorDimensions {

	dims := NewDims()
	dims.floor = SetDimensions(int(winDims.Height), int(winDims.Width), len(floors))
	dims.positions = SetCarPositions(floors, dims.floor)

	return dims
}

func Background(dims fyne.Size, floorDims FloorDimensions, levels []*Level) *fyne.Container {

	cont := container.NewWithoutLayout()

	// Background rectangle
	//
	backgroundbox := canvas.NewRectangle(DARK)
	backgroundbox.Resize(dims)
	backgroundbox.SetMinSize(dims)
	cont.Add(backgroundbox)
	

	// Floors
	//
	for index, level := range levels {

		height := index*floorDims.floorHeight + floorDims.bottomLevel
		flObj := CreateFloorObject(ycoordOffset, height, floorDims, level.Front, level.Rear)

		for _, obj := range flObj {
			cont.Add(obj)
		}
	}

	return cont
}


func CreateFloorObject(yOff int, yLevel int, floorDims FloorDimensions, front bool, rear bool) []fyne.CanvasObject {

	objs := []fyne.CanvasObject{}
	position := fyne.Position{X: 0, Y: float32(yLevel)}

	if front {

		line, pos := CreateALine(BLACK, yOff, position, 2, floorDims.hallLength)
		objs = append(objs, line)
		position = pos

		line, pos = CreateALine(RED, yOff, position, 4, floorDims.lobbyLength)
		objs = append(objs, line)
		position = pos

	} else {
		position.X = position.X + float32(floorDims.hallLength+floorDims.lobbyLength)
	}

	line, pos := CreateALine(GREY, yOff, position, 2, floorDims.carLength)
	objs = append(objs, line)
	position = pos

	if rear {
		line, pos := CreateALine(RED, yOff, position, 4, floorDims.lobbyLength)
		objs = append(objs, line)
		position = pos

		line, pos = CreateALine(BLACK, yOff, position, 2, floorDims.hallLength)
		objs = append(objs, line)
		position = pos
	}

	return objs
}

// Position1 == top-left  Position2 == bottom-right
func CreateALine(color color.RGBA, yOff int, position fyne.Position, width int, length int) (*canvas.Line, fyne.Position) {
	anyLine := canvas.NewLine(color)
	anyLine.StrokeWidth = float32(width)

	anyLine.Position1 = flipVertical(yOff, position)
	position.X = position.X + float32(length)
	anyLine.Position2 = flipVertical(yOff, position)

	return anyLine, position
}

func flipVertical(height int, pos fyne.Position) fyne.Position {
	return fyne.Position{
		X: pos.X,
		Y: float32(height) - pos.Y,
	}
}
