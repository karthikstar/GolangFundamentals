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

type Motorcycle string
type Car string
type Truck string

type VehicleDirecter interface {
	directVehicle()
}

// Motorcycle type implements the VehicleDirecter interface
func (m Motorcycle) directVehicle() {
	fmt.Println("Direct towards small lift")
}

// Car type implements the VehicleDirecter interface
func (c Car) directVehicle() {
	fmt.Println("Direct towards standard lift")
}

// Truck type implements the VehicleDirecter interface
func (t Truck) directVehicle() {
	fmt.Println("Direct towards large lift")
}

// this function takes a slice of the VehicleDirecter interface as a parameter
// as seen in line 58, we are passing in a slice thatt contains 3 diff types - motorcycle, car, truck to this function
// this function directs Vhicles by calling the directVehicle method of the CORRESPONDING TYPE
func directVehicles(vehicles []VehicleDirecter) {
	fmt.Println("Directing vehicles...")
	for i := 0; i < len(vehicles); i++ {
		vehicle := vehicles[i]
		vehicle.directVehicle()
		fmt.Printf("Vehicle Information : %v\n", vehicle)
	}
	fmt.Println()
}
func main() {
	vehicles := []VehicleDirecter{Motorcycle("Yahama Motorcycle"), Car("Ferrari car"), Truck("Toto Truck")}
	directVehicles(vehicles)
}
