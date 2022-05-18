//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

//  - Security tags have two states: active (true) and inactive (false)
const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

//* Create a structure to store items and their security tag state
type Item struct {
	name string
	tag  SecurityTag
}

//* Create functions to activate and deactivate security tags using pointers
func activate(tag *SecurityTag) {
	*tag = Active
}
func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkOut(items []Item) { //takes in a slice/array of Item datatype
	fmt.Println("checking out..")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func main() {
	//  - Create at least 4 items, all with active security tags
	shirt := Item{"Shirt", Active}
	pants := Item{"Pants", Active}
	purse := Item{"Purse", Active}
	watch := Item{"Watch", Active}
	//  - Store them in a slice or array
	items := []Item{shirt, pants, purse, watch}
	fmt.Println("initial", items)
	//  - Deactivate any one security tag in the array/slice
	deactivate(&items[0].tag)
	fmt.Println("item 0 deactivated", items)
	checkOut(items)
	fmt.Println("after checkout", items)

}
