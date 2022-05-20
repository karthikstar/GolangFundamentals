package main //whenever we have a main package, that will be the pkg that launches the whole program. it will refer to the other packages

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display") // all we need to do is to refer to the last name in the package path abv e.g display. and it will automatically put in the required imports
	msg.Exciting("an exciting message")
}

// we have display package, msg package, and main package
//msg package has 2 files - every file that uses package msg at the top are smashed together into 1 single package
