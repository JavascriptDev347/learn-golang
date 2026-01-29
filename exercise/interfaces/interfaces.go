//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Mechanic interface {
	gotToLift()
}

type SmallVehicle string
type StandardVehicle string
type LargeVehicle string

func (s SmallVehicle) gotToLift() {
	fmt.Printf("%s is small lift\n", s)
}

func (st StandardVehicle) gotToLift() {
	fmt.Printf("%s is standard lift\n", st)
}

func (l LargeVehicle) gotToLift() {
	fmt.Printf("%s is large lift\n", l)
}

func directVehicles(vehicles []Mechanic) {
	for _, v := range vehicles {
		v.gotToLift()
	}
}

func main() {
	vehicles := []Mechanic{SmallVehicle("Yamaha R1"), StandardVehicle("Toyota Camry"), LargeVehicle("Ford F-150")}
	directVehicles(vehicles)
}
