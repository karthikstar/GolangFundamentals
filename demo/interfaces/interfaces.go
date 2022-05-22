package main

// interface is for dish preparation
import "fmt"

type Preparer interface {
	// naming convention for interfaces is to have 'er' at the end - not compuslory tho
	PrepareDish()
}

// 2 diff dishes
// type aliases
type Chicken string
type Salad string

// receiver function that takes in a value
//preparing chicken
func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

//preparing salad
func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

// fn to prepare multiple dishes
// takes in a slice of preparer
func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish :%v--\n", dish)
		dish.PrepareDish() // implemented on both chicken and salad
	}
	fmt.Println()
}

func main() {
	// we are creating a new Slice that has Preparer interface items
	// everything within this slice has to implement the Preparer Interface
	dishes := []Preparer{Chicken("Chicken Wings"), Salad("Iceberg Salad")}
	prepareDishes(dishes)
}
