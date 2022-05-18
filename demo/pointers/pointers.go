package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	//counter is a pointer to a Counter
	// when we have  a pter to a structure we dunnid to dereference it
	// however for normal pters to integer etc we need to dereference using *
	counter.hits += 1
	fmt.Println("Counter", counter)
}
func replace(old *string, new string, counter *Counter) {
	*old = new // new string getting copied over to old
	increment(counter)
}

func main() {
	counter := Counter{}
	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)
	replace(&hello, "Hi", &counter) // replace hello with Hi w the help of ptrs
	fmt.Println(hello, world)

	//create a slice
	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)

}
