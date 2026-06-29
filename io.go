package main

import (
	"log"
	"os"
	"encoding/json"
)

var Levels []*Level


func ReadLevels(fname string)  error {

	file, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	var levels []*Level
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&levels); err != nil {
		return err
	}

	Levels = levels
	
	return nil
}

func ShowLevels(levels []*Level) {
	log.Println("Number of floors", len(levels))
	for _, level := range levels {
		log.Println("lev", level)
	}
}
