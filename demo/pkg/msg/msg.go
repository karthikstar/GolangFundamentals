package msg

import "coursecontent/demo/pkg/display"

// import "display" - we CANT just do this

//import "coursecontent/demo/pkg/display" // we need check go.mod for the root package

func Hi() {
	display.Display("Hi")
}
