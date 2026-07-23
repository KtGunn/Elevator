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



///////////////////////////////////////////////////////////
// LEVEL ('upper case!')
//
type Level struct {
  Number int32
	Front  bool
	Rear   bool
}


//////////////////////////////////////////////////////////////
// Elevator
//
type Elevator struct {
	bank string
	image *fyne.Container

	car *Car

	dimensions ElevatorDimensions
	levels []*Level
}

func NewElevator(bank string) *Elevator {
	return &Elevator{
		bank: bank,
	}
}


func (e *Elevator) Dimension(dims fyne.Size, floors int) {

	e.dimensions = ElevatorDimensions{}
	e.dimensions = e.dimensions.Dimensions(
		int(dims.Height),
		int(dims.Width), floors,
	)
}


func (e *Elevator) Levels(landings []*Landing) {
	
	e.levels = make([]*Level, 0)
	
	landIndex := 0
	for pi := 0; pi < e.dimensions.floorsCount; pi++ {
		
		done := false
		for n := landIndex; !done ;landIndex++ {
			landing := landings[n]
			
			if int(landing.Floor) == pi {
				level := &Level{
					Number: int32(landing.Floor),
					Front: landing.Door == 0 || landing.Door == 2,
					Rear:  landing.Door == 2 || landing.Door == 1,
				}
				
				e.levels = append(e.levels, level)
				done = true  // advance landIndex
				
			} else {
				e.levels = append(e.levels, nil)
				break   // don't advance landIndex
			}
		}
	}
}


func (e *Elevator) Car(car string) {
	e.car = CreateCar(car, e.dimensions.car)
}

func (e *Elevator) SetCar(floor int) {
	//	e.image.Add(e.car.image)
}

// ElevatorDims
//
func ElevatorDims(winDims fyne.Size, floors int) ElevatorDimensions {

	return ElevatorDimensions{}
}


func (e *Elevator) Place(floor int) {
	log.Println("@Place ElevatorCar: (global)", yOffset)
}


// Image
//
func (e *Elevator) Image(win fyne.Size) {

	e.image = container.NewWithoutLayout()

	// Background rectangle
	imagebox := canvas.NewRectangle(DARK)
	imagebox.Resize(win)
	imagebox.SetMinSize(win)

	e.image.Add(imagebox)


	dims := e.dimensions
	log.Println("@e.Image dims", e.dimensions)

	// Floors
	//
	log.Println(" ready for levels ...")
	for index, level := range e.levels {
		log.Println("index", index)

		if level == nil {
			log.Println("NIL")
			continue  // a floor this elevator does not service
		}

		height := index*dims.floor.floorHeight + dims.floor.bottomLevel
		flObj := CreateFloorObject(int(win.Height), height, dims, level.Front, level.Rear)

		for _, obj := range flObj {
			log.Println(" ... obj added")
			e.image.Add(obj)
		}
	}

	return
}


func CreateFloorObject(yOff int, yLevel int, dims ElevatorDimensions, front bool, rear bool) []fyne.CanvasObject {

	objs := []fyne.CanvasObject{}
	position := fyne.Position{X: 0, Y: float32(yLevel)}

	if front {

		line, pos := CreateALine(BLACK, yOff, position, 2, dims.floor.hallLength)
		objs = append(objs, line)
		position = pos

		line, pos = CreateALine(RED, yOff, position, 4, dims.floor.lobbyLength)
		objs = append(objs, line)
		position = pos

	} else {
		position.X = position.X + float32(dims.floor.hallLength+dims.floor.lobbyLength)
	}

	line, pos := CreateALine(GREY, yOff, position, 2, dims.car.carLength)
	objs = append(objs, line)
	position = pos

	if rear {
		line, pos := CreateALine(RED, yOff, position, 4, dims.floor.lobbyLength)
		objs = append(objs, line)
		position = pos

		line, pos = CreateALine(BLACK, yOff, position, 2, dims.floor.hallLength)
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


// flipVertical
//   this is a quick coordinates transformation.
//   Screen coords start at top-left, x  positive right and y position down.
//   The application sets origin at bottom left, x  positive right and y position UP.
//
func flipVertical(height int, pos fyne.Position) fyne.Position {
	return fyne.Position{
		X: pos.X,
		Y: float32(height) - pos.Y,
	}
}
