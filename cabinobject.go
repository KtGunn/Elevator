package main

import (
	"fmt"
	"fyne.io/fyne/v2"
)


type CabinObject struct {
	elevator ElevatorCabin
	image *fyne.Container

	bank string
	cabin string
}
func NewCabinObject (bank string, cabin string) CabinObject {
	return CabinObject{
		bank: bank,
		cabin: cabin,
	}
}


// GetCabinObject
//  Returns the CabinObject of the bank OR cabin name
//
func GetCabinObject(bank string, car string) (CabinObject, error) {

	for _, co := range CabinObjects {
		if bank == co.bank {
			return co, nil
		}
		if car == co.cabin {
			return co, nil
		}
	}

	return CabinObject{}, fmt.Errorf("cabin object not found")
}


// CarFromName
//  Returns the *Car object of the named cabin
//
func CarFromName(name string) *Car {
	for _, co := range CabinObjects {
		if name == co.cabin {
			return co.elevator.car
		}
	}
	return nil
}



