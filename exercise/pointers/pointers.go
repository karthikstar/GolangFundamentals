//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
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

type secTag struct {
	item   string
	active bool // true/false
}

func activateTag(tag *secTag) {
	tag.active = true
}

func deactivateTag(tag *secTag) {
	tag.active = false
}

func checkout(array []secTag) {
	for i := range array { //deactivate all the tags
		array[i].active = false
	}
}

func main() {
	tags := []secTag{
		{"pen", true},
		{"pencil", true},
		{"phone", true},
		{"laptop", true},
	}
	fmt.Println("initial", tags)

	deactivateTag(&tags[1])
	fmt.Println("after deactivation of 2nd tag", tags)
	checkout(tags)
	fmt.Println("aft deactivating all", tags)

}
