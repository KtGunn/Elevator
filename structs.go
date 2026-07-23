package main

import (
	"fyne.io/fyne/v2"
)


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




type FloorObject struct {
	shape    fyne.CanvasObject
	size     fyne.Size
	position fyne.Position
}
