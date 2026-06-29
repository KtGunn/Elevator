package main

import (
	"log"
	"fyne.io/fyne/v2"
)


func main() {
	var err error

	err = ReadLevels("levels.json")
	if err != nil {
		log.Fatal("could not read json file: ", err)
	}
	
	windowSize :=fyne.Size{
		Width: 240,
		Height: 280,
	}

	CreateAppInstance(windowSize, Levels)
	CreateElevatorCabin(windowSize, Levels)

	return
}



