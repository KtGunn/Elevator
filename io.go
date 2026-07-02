package main

import (
	"log"
	"os"
	"encoding/json"
)

var Banks []*Bank


func ReadBanks(fname string)  error {

	file, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	var banks []*Bank

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&banks); err != nil {
		return err
	}

	log.Println("  banks --- ")
	for _, b := range banks {
		log.Println("bank name", b.Name)
		for _, c := range b.Cars {
			log.Println("  cab name", c.Name)
			for _, l := range c.Landings {
				log.Println("    landing:", l.Floor, l.Door, l.Label)
			}
		}
	}

	Banks = banks
	
	return nil
}

func ShowLevels(levels []*Level) {
	log.Println("Number of floors", len(levels))
	for _, level := range levels {
		log.Println("lev", level)
	}
}
