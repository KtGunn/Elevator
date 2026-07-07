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
