package main // specifies that we have a main package

import "fmt" // we are importing the fmt - format library

func main() { // first function that is ran
	var myName = "Karthik"
	fmt.Println("my name is", myName, myName)

	//create a variable named name and set its type to string.
	var name string = "Kathy" // type annotation
	fmt.Println("name = ", name)

	userName := "admin"
	fmt.Println("username=", userName)

	var sum int
	fmt.Println("The sum is", sum)

	//compound assignment
	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)

	//ok, error idiom
	part2, other := 2, 0
	fmt.Println("part 2 is", part2, "other is", other) // other can be reused twice

	sum = part1 + part2
	fmt.Println("sum =", sum)

	//block assignment
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "hello", "world", "!"
	//we wont have ! as it gets ignored due to the _

	fmt.Println(word1, word2)

}
