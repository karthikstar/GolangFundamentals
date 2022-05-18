//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printSlice(slice []Part) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println()
}

func main() {
	assemblyLine := []Part{"kit", "spanner", "nut"} // 3 parts
	fmt.Println("Part 1")
	printSlice(assemblyLine)

	assemblyLine = append(assemblyLine, "ruler", "eraser") // 5 parts
	fmt.Println("Part 2")
	printSlice(assemblyLine)

	assemblyLine = assemblyLine[3:] // 2 parts
	fmt.Println("Part 3")
	printSlice(assemblyLine)
}
