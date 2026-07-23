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

	ShowBanks(banks)
	Banks = banks
	
	return nil
}

func NumberOfFloors() int {
	min, max := MinMaxLandings(Banks)
	log.Println("NumberOfFloors","min", min, "max", max)
	return max-min+1
}

// MinMaxLandings
//  returns the lowest + highest floors
//  specified across all banks
//
func MinMaxLandings(banks []*Bank) (int, int) {
	minFloor := 1000
	maxFloor := -1
	
	for _, b := range banks {
		for _, c := range b.Cars {
			for _, l := range c.Landings {
				if int(l.Floor) < minFloor {
					minFloor = int(l.Floor)
				}
				if int(l.Floor) > maxFloor {
					maxFloor = int(l.Floor)
				}
			}
		}
	}
	return minFloor, maxFloor
}

func ShowBanks(banks []*Bank) {
	log.Println("Banks...")

	for _, b := range banks {
		log.Println("bank name", b.Name)
		for _, c := range b.Cars {
			log.Println("  cab name", c.Name)
			for _, l := range c.Landings {
				log.Println("    landing:", l.Floor, l.Door, l.Label)
			}
		}
	}
	min, max := MinMaxLandings(banks)
	log.Println("Min floor", min, "max floor", max)
}

