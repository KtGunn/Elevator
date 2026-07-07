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
	WIDTH_CAR   int = 10
	WIDTH_LOBBY int = 12
	WIDTH_HALL  int = 18

	HEIGHT_BOX_MAX int = 50
	WIDTH_BOX_MAX  int = 50

	VERTICAL_PIXEL_MARGIN int = 2
)




/////////////////////////////////////////////////////////////////////
// ACTUAL FLOOR DIMENSIONS in pixels.
//
type RobotDimensions struct {
	bodyHeight int
	bodyWidth  int
	wheelDia   int
}

type FloorDimensions struct {
	hallLength   int
	lobbyLength  int
	floors       int
	floorHeight  int
	bottomLevel int
}

// Cabin & Doors
//
//    ______________
//    |_         __|
//    | |        | |
//    | |        | |
//    | |        | |
//    |_|________|_|
//
type CarDimensions struct {
	carLength int
	boxHeight int
}



//
type ElevatorDimensions struct {
	floor     FloorDimensions
	car       CarDimensions
	positions   []CarPosition
}




func SetDimensions(winHeight int, winWidth int, floors int) (FloorDimensions, CarDimensions) {

	var floorDims FloorDimensions = FloorDimensions{
		floors: floors,
	}

	var carDims CarDimensions = CarDimensions{}
	
	floorDims.floorHeight, carDims.boxHeight, floorDims.bottomLevel =
		FloorAndCabHeights(winHeight, winWidth, floorDims.floors)

	carDims.carLength, floorDims.lobbyLength, floorDims.hallLength =
		AllocateDimensions(int(winHeight), int(winWidth))

	return floorDims, carDims
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



// SetCarPositions
// The function returns a slice of CarPositions for each level
// The car positions are used to place the car in the display
// given the level or floor it is on.
//
// -- NOTE mauybe I should be using a map instead of a slice?
//
func SetCarPositions(levels []*Level, floorDims FloorDimensions) []CarPosition {

	carPositions := []CarPosition{}
	xCoord := floorDims.hallLength + floorDims.lobbyLength
	
	for _, level := range levels {
		number := int(level.Number)
		yCoord := number*floorDims.floorHeight + floorDims.bottomLevel
		carPositions = append(carPositions, CarPosition{
			level: number,
			xPixCoord: xCoord,
			yPixCoord: yCoord,
		})
	}

	return carPositions
}



// CabinToLevels
//  Translates Landings from input json file into the Level structure
//  used to draw the elevator cabin.
//
func CabinToLevels(landings []*Landing) []*Level {

	var levels []*Level

	for _, landing := range landings {
		level := &Level{
			Number: int32(landing.Floor),
			Front: landing.Door == 0 || landing.Door == 2,
			Rear:  landing.Door == 2 || landing.Door == 1,
		}
		levels = append(levels, level)
	}

	return levels
}
