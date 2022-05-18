//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}
type Rectangle struct {
	a Coordinate //top left
	b Coordinate //bottom right
}

func width(rect Rectangle) int {
	return (rect.b.x - rect.a.x) // width of rectangle
}

func length(rect Rectangle) int {
	return (rect.a.y - rect.b.y) // length of rectangle
}

func calcArea(rect Rectangle) int {
	return length(rect) * width(rect)
}

func calcPerimeter(rect Rectangle) int {
	return 2*length(rect) + 2*width(rect)
}

func printInfo(rect Rectangle) {
	fmt.Println("Area is", calcArea(rect))
	fmt.Println("Perimeter is", calcPerimeter(rect))
}
func main() {
	rect1 := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	printInfo(rect1)
	// to douple the size
	rect1.a.y *= 2
	rect1.b.x *= 2
	printInfo(rect1)
}
