//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function parameter and displays a greeting to that person.
func greet(name string) {
	println("Hello there,", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func callOwnself() string {
	return "wassuppp"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func sum(a, b, c int) int {
	return a + b + c
}

//* Write a function that returns any number
func returnAnyNo() int {
	return 2
}

//* Write a function that returns any two numbers
func returnAny2Nos() (int, int) {
	return 1, 3
}

//* Add three numbers together using any combination of the existing functions.
func addThreeNos(a, b, c int) int {
	return a + b + c
}

//  * Print the result

//* Call every function at least once
func main() {

	greet("Karthik")
	fmt.Println("now we will call the function", callOwnself())
	fmt.Println(returnAny2Nos())
	fmt.Println(returnAnyNo())
	fmt.Println("sum of 3 nos is", addThreeNos(1, 2, 3))
	fmt.Println("sum of 3 nos is", sum(1, 2, 3))

	a, b := returnAny2Nos()

	answer := addThreeNos(returnAnyNo(), a, b)
	fmt.Println(answer)
}
