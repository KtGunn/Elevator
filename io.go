package main

import (
	"log"
	"os"
	"encoding/json"
)


func ReadLevels(fname string) ([]*Level, error) {

	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var levels []*Level
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&levels); err != nil {
		return nil, err
	}

	return levels, nil
}

func ShowLevels(levels []*Level) {
	log.Println("Number of floors", len(levels))
	for _, level := range levels {
		log.Println("lev", level)
	}
}
