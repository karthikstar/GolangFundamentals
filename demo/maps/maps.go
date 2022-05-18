package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11 // adding a key value pair
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1 // will assign a 0 (default) thne +1
	shoppingList["eggs"] += 1  // will give 11 + 1 = 12
	fmt.Println(shoppingList)
	delete(shoppingList, "milk")
	fmt.Println("milk deelted", shoppingList)
	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yup", cereal, "boxes")
	}

	totalItems := 0
	for _, amount := range shoppingList { // key, value - amt
		totalItems += amount
	}
	fmt.Println("total items: ", totalItems)

}

//maps store data in key vaue pairs. if we know the key, easy to acess the data
// myMap := make(map[string]int)
// map[data type for key] data type for values

// with maps we dunnid to specify no of elements
// size of map is dynamic, so we dont need to list any size, just list the numbers would do
