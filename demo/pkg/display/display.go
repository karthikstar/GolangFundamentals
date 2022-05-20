package display //this is part of the display package. not part of main package
//we cant use the main fn here

import "fmt"

func Display(msg string) { // NOTE: when we have a capital letter for fn at start, it will be exported. this fn will be avail outside the display fn
	fmt.Println(msg)
}

//func hello(msg string) { // this wont be accessible in the main.go file
//	fmt.Println(msg)
//}

// thus, public or private fns in go just depends on whether we capitalise the first letter of the fn
