package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

//normal function
func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

// receiver function
func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println("Initial:", lot)
	fmt.Println(lot)
	lot.occupySpace(1)                  // occupy space 1
	occupySpace(&lot, 2)                // occupy space 2
	fmt.Println("After occupied:", lot) // spaces 1 and 2 are true. 3 and 4 are false
	lot.vacateSpace(2)
	fmt.Println("After vacate:", lot) // space 1 is true, 2 3 4 is false
}
