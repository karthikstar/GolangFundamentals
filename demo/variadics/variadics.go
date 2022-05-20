package main

import "fmt"

// there will be situations where we write functions and we dont know how many parameters it will have
//variadics is  a way to write a fn which accepts any no of parameters

//whenever we use ...type, it means the parameter is going to be a slice
// if we have a slice, we need to iterate thru it

func sum(nums ...int) int { //3 dots indicates its a variadics. we have nums var of type integer. we can have any no of integers listed
	//e.g 1,2 .. 100
	// when we have variadics, nums variable is actl a slice of integers. we need to iterate thru them
	sum := 0
	for _, n := range nums { // ignore index, we just need the no n.
		sum += n
	}
	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	all := append(a, b...)
	answer := sum(all...) // when we use ... with the slice, its going to expand to each of these integers individually.
	// the function signature takes in any no of integers, 1,2 or 30
	fmt.Println(answer)

	//when we have slice e.g ans with 3 dots ... what we are essentially doing is
	answer = sum(1, 2, 3, 4, 5, 6)
	fmt.Println(answer)
}

// when we need a variable no of parameters, then we shld use variadics as it will allow us to have both slices and as many manually entered nos as we want
