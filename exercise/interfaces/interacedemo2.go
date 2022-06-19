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

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

// implement preparer interface on chicken

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v-- \n", dish)
		dish.PrepareDish() // call prepare dish fn which has been implemeneted on both chicken n salad
	}
	fmt.Println()
}

func main() {
	// creating a slice which has the preparer inteface
	// everything within this slice implements the interface
	dishes := []Preparer{Chicken("Chicken Wings"), Salad("Iceberg Salad")}
	prepareDishes(dishes)
}
