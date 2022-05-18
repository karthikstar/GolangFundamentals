package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	//slices and underlying array can be created at the same time as shown below
	route := []string{"Grocery", "Department Store", "Salon"} // making the string slice, of 3 items
	printSlice("Route 1", route)
	route = append(route, "Home") // we are reassinging route, and appending home to it. now route has 4 items
	printSlice("Route 2", route)
	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")
	route = route[2:]
	printSlice("Remaining", route)
}
