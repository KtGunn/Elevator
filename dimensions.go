package main

import (
	"log"
)


/////////////////////////////////////////////////////////////////////
// RELATIVE FLOOR DIMENSIONS, not actual pixel dimensions.
//
//                       |
//    (hallway)    (lobby) | (cabin)  C (center
//   ............|_________|_________ L  line)
//
//   size ratios: hallway(10) lobby(10) cabin(6)
//
const (
	WIDTH_CAR   int = 10  // Relative dimensions
	WIDTH_LOBBY int = 12
	WIDTH_HALL  int = 18
	
	HEIGHT_BOX_MAX int = 50 // Absolute pixel dimensions
	WIDTH_BOX_MAX  int = 50

	VERTICAL_PIXEL_MARGIN int = 2
)



/////////////////////////////////////////////////////////////////////
// ACTUAL FLOOR DIMENSIONS in pixels.
//
type FloorDimensions struct {
	hallLength   int
	lobbyLength  int
	carVoid      int

	floorHeight  int
	bottomLevel int
}



/////////////////////////////////////////////////////////////////////
//  FLOOR POSITIONS
//
//                  fOut               rOut
//   o-----x-----o-x--x--x-o|__x__|o-x--x--x-o----x----0
//      fHall    fLby   fAtC  InC   rAtC  rLby   rHall
//   
const (
	FRONT_SIDE int = 1
	REAR_SIDE  int = 2
)

func (d FloorDimensions) xPosition(side int, pcol int) int {
	
	var pos int
	
	switch pcol {
	case PCOL_RESERVE:
		pos = d.hallLength/2
	case PCOL_LOBBY:
		pos = d.hallLength + 2
	case PCOL_OUTCAR:
		pos = d.hallLength + d.lobbyLength/2
	case PCOL_ATCAR:
		pos = d.hallLength + d.lobbyLength-2
	case PCOL_INCAR:
		pos = d.hallLength + d.lobbyLength + d.carVoid/2
	case PCOL_DONE:
		pos = 5
	default:
		return 0
	}

	switch side {

	case FRONT_SIDE:
		return pos

	case REAR_SIDE:
		return 2*(d.hallLength + d.lobbyLength) + d.carVoid - pos
	}

	return 0
}

type RobotDimensions struct {
	bodyHeight int
	bodyWidth  int
	wheelDia   int
}

// Cabin & Doors
//
//    |< carLength >|
//    _______________
//    |_         __|  ^
//    | |        | |  |
//    | |        | |  carheight
//    | |        | |
//    |_|________|_|  |
//                    ^
type CarDimensions struct {
	carLength int
	boxHeight int
}



//
type ElevatorDimensions struct {
	floor     FloorDimensions
	car       CarDimensions
	floorsCount  int
}



func (d ElevatorDimensions) xyPosition(floor int, side int, pcol int) (int, int) {
	x := d.floor.xPosition(side, pcol)
	y := d.floor.bottomLevel + floor * d.floor.floorHeight
	return x,y
}


func (d ElevatorDimensions) Dimensions(winHeight int, winWidth int, floors int) ElevatorDimensions {

	d.floor = FloorDimensions{}
	d.car = CarDimensions{}
	
	d.floor.floorHeight, d.car.boxHeight, d.floor.bottomLevel =
		FloorAndCabHeights(winHeight, winWidth, floors)

	d.car.carLength, d.floor.lobbyLength, d.floor.hallLength =
		AllocateDimensions(int(winHeight), int(winWidth))

	d.floor.carVoid = d.car.carLength
	d.floorsCount = floors

	return d
}




func FloorAndCabHeights(overallHeight int, overallWidth int, floors int) (int, int, int) {

	var floorHeight int
	var boxHeight int
	
	verticalMargin := 2 * VERTICAL_PIXEL_MARGIN
	
	floorHeight = (overallHeight-verticalMargin) / floors

	if floorHeight > HEIGHT_BOX_MAX {
		boxHeight = HEIGHT_BOX_MAX
	} else {
		boxHeight = floorHeight - 2
	}

	return floorHeight, boxHeight, VERTICAL_PIXEL_MARGIN
}


func AllocateDimensions(overallHeight int, overallWidth int) (int, int, int) {
	log.Println(overallWidth)
	var boxWidth int
	var cabinWidth float32
	
	units := 2*(WIDTH_HALL + WIDTH_LOBBY) + WIDTH_CAR
	unitWidth := float32(overallWidth)/float32(units)
	
	cabinWidth = float32(WIDTH_CAR) * unitWidth
	
	if int(cabinWidth) > WIDTH_BOX_MAX {
		boxWidth = WIDTH_BOX_MAX
	} else {
		boxWidth = int(cabinWidth)
	}

	floatUnits := float32((overallWidth - boxWidth)) / float32((2*(WIDTH_HALL + WIDTH_LOBBY)))
	return boxWidth, int(floatUnits*float32(WIDTH_LOBBY)), int(floatUnits*float32(WIDTH_HALL))
}



