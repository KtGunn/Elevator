package main

import (
	//"log"
)



func SetDimensions(graphicsHeight int, graphicsWidth int, floors int) FloorDimensions {

	var floorDims FloorDimensions = FloorDimensions{
		floors: floors,
	}

	floorDims.floorHeight, floorDims.boxHeight, floorDims.bottomLevel =
		FloorAndCabHeights(graphicsHeight, graphicsWidth, floorDims.floors)

	floorDims.carLength, floorDims.lobbyLength, floorDims.hallLength =
		AllocateDimensions(int(graphicsHeight), int(graphicsWidth))

	return floorDims
}



var verticalPixelMargin int = 2

func FloorAndCabHeights(overallHeight int, overallWidth int, floors int) (int, int, int) {

	var floorHeight int
	var boxHeight int
	
	verticalMargin := 2 * verticalPixelMargin
	
	floorHeight = (overallHeight-verticalMargin) / floors

	if floorHeight > 50 {
		boxHeight = 50
	} else {
		boxHeight = floorHeight - 2
	}

	return floorHeight, boxHeight, verticalPixelMargin
}


func AllocateDimensions(overallHeight int, overallWidth int) (int, int, int) {

	var boxWidth int
	var cabinWidth float32

	units := 2*(10+10)+6
	unitWidth := float32(overallWidth)/float32(units)
	
	if false {
		if cabinWidth > 50 {
			boxWidth = 50
		} else {
			boxWidth = int(cabinWidth)
		}
		units = (overallWidth - boxWidth) /(2*(10+10))
	}
	
	return int(6*unitWidth), int(10*unitWidth), int(10*unitWidth)
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
	
	for _, level := range Levels {
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
