package main

import (
	"fyne.io/fyne/v2"
)


///////////////////////////////////////////////////////////
// LEVEL ('upper case!')
//
type Level struct {
  Number int32
	Front  bool
	Rear   bool
}


///////////////////////////////////////////////////////////
// I/O structures
// These are used with a json file for input
//
type Bank struct {
	Name string   `json:name`
	Cars []*Cabin `json:Cabins`
}
type Cabin struct {
	Name string         `json:name`
	Landings []*Landing `json:landing`
}
type Landing struct {
	Floor int8  `json:"Floor"`
	Door  int8  `json:"Door"`
	Label string `json:"Label"`
}




// FLOOR DIMENSIONS
//
/*                       |
    (hallway)    (lobby) | (cabin)  C (center
   ............|_________|_________ L  line)

   size ratios: hallway(10) lobby(10) cabin(6)
*/


// Cabin & Doors
/*
    ______________
    |_         __|
    | |        | |
    | |        | |
    | |        | |
    |_|________|_|
*/


//
type ElevatorDimensions struct {
	floor       FloorDimensions
	positions   []CarPosition
}

func NewDims() ElevatorDimensions {
	return ElevatorDimensions{}
}

type FloorDimensions struct {
	hallLength   int
	lobbyLength  int
	carLength    int
	floorHeight  int
	boxHeight    int
	floors       int
	bottomLevel int
}

type CarPosition struct {
	level int
	xPixCoord int
	yPixCoord int
}

type FloorObject struct {
	shape    fyne.CanvasObject
	size     fyne.Size
	position fyne.Position
}
