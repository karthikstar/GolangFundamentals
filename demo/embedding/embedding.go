package main

import "fmt"

// simulates a shipping warehouse
// we will be using type embedding

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

// use iota enumeration pattern on beltsize and shipping

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

// creating some interfaces
type Conveyor interface {
	Convey() BeltSize // return BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	// embedding 2 interfaces
	// requires both to be implemented
	// both Convey() and Ship() has to be implemented
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

// spamMail has Ship and convey implemented
func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

// toxic waste only has ship implemented so we are unable to use it with the automate function
func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	mail := SpamMail{40000}
	automate(&mail)

	//	waste := ToxicWaste{300}
	//	automate(&waste) // this not possible as
	// we didnt implement BOTH interfaces
}
