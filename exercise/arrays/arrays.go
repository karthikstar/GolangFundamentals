//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	price int
	name  string
}

func totalNoCost(list [4]Product) {
	var cost, totalNoOfItems int
	for i := 0; i < len(list); i++ {
		cost += list[i].price
		if list[i].name != "" {
			totalNoOfItems++
		}
	}
	fmt.Println("Last item on list", list[totalNoOfItems-1])
	fmt.Println("Total items", totalNoOfItems)
	fmt.Println("Total cost", cost)
}

func main() {
	//creating shopping list with 4 products
	shoppingList := [4]Product{
		{100, "eggs"},
		{20, "bread"},
		{50, "chicken"}, //default value for 4th item will be {0,""}
	}
	totalNoCost(shoppingList)
	shoppingList[3] = Product{20, "Cake"}
	totalNoCost(shoppingList)
}
