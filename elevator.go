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
	car        *fyne.Container
	floors     *fyne.Container
	background *fyne.Container
}

func NewElevatorCabin() ElevatorCabin {
	return ElevatorCabin{}
}



//////////////////////////////////////////////////////////////
// CreateElevatorCabin
// creates all objects needed to render an elevator cabin
//
func CreateElevatorCabin(dims fyne.Size, levels []*Level) ElevatorCabin {

	elevDims := NewDims()
	elevDims.floor, elevDims.positions = ElevatorDims(dims, elevDims.floor,levels)

	newCabin := NewElevatorCabin()
	newCabin.floors = FloorsContainer(dims, elevDims.floor, levels)
	newCabin.background = Background(dims)
	newCabin.car = NewCar(elevDims.floor)

	return newCabin
}

func Background(dims fyne.Size) *fyne.Container {

	backgroundbox := canvas.NewRectangle(DARK)
	backgroundbox.Resize(dims)

	return container.NewWithoutLayout(backgroundbox)
}

func ElevatorDims(winDims fyne.Size, floorDims FloorDimensions, floors []*Level) (FloorDimensions, []CarPosition){

	floorDims = SetDimensions(int(winDims.Height), int(winDims.Width), len(floors))
	carPositions := SetCarPositions(floors, floorDims)

	return floorDims, carPositions
}


func FloorsContainer(dims fyne.Size, floorDims FloorDimensions, levels []*Level) *fyne.Container {
	vBox := CreateFloors(int(dims.Height), floorDims, levels)
	vBox.Resize(fyne.NewSize(dims.Width, dims.Height))
	return vBox
}

func CreateFloors(ycoordOffset int, floorDims FloorDimensions, levels []*Level) *fyne.Container {
	log.Println("Creating floors.")

	vbox := container.NewWithoutLayout()

	for index, level := range levels {

		height := index*floorDims.floorHeight + floorDims.bottomLevel
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
