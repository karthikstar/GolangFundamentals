//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func main() {

	// age := 10
	// switch age {
	// case 0:
	// 	fmt.Println("newborn")
	// case 1, 2, 3:
	// 	fmt.Println("toddler")
	// case 4, 5, 6, 7, 8, 9, 10, 11, 12:
	// 	fmt.Println("child")
	// case 13, 14, 15, 16, 17:
	// 	fmt.Println("teenager")
	// default:
	// 	fmt.Println("adult")
	// }

	//better ver
	switch age := 14; { // remember need include ; at the end so that we can create and assign a variable right there
	case age == 0:
		fmt.Println("newborn")
	case age >= 1 && age <= 3:
		fmt.Println("toddler")
	case age < 13:
		fmt.Println("child")
	case age < 18:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")
	}

}
