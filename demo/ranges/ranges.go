package main

import "fmt"

func main() {
	//when we use ramge keyword, its going to automatically create a iterator for us
	// one advantage is we dont need to calculate length in order to traverse
	slice := []string{"Hello", "world", "!"}
	for i, element := range slice { // when we use range, we are going to get 2 things from the slice. i is index, elemenet is hellow/world/!
		fmt.Println(i, element)
		// note: we cannot use for loop to traverse thru individual letters

		//however, we can use range keyword to iterate thru individual letters
		for _, ch := range element {
			fmt.Printf("  %q\n", ch) // each of these is a rune
		}
	}
}
