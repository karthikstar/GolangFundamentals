package main

import "fmt"

const (
	Low = iota
	Medium
	High
)

// we are creating a priority queue - allow us to add items of diff priorities - take item of highest priority out first

type PriorityQueue[P comparable, V any] struct { // generic type T must be compable and generic type V can be any
	items      map[P][]V // keys are going to be P, and the values are going to be a slice of V
	priorities []P       // slice of P
}

// rec function add a item to priroirty queue
func (pq *PriorityQueue[P, V]) Add(priority P, value V) { // since this is a rec function we dunnid to specify the generic constraints again so can just use P and V
	pq.items[priority] = append(pq.items[priority], value) // adding a extra value to existing items at a priority. appends at the end
}

// pulling out the next item in the priority queue - returns v - value and bool - whthr something was actl returned
func (pq *PriorityQueue[P, V]) Next() (V, bool) {
	for i := 0; i < len(pq.priorities); i++ {
		priority := pq.priorities[i]
		items := pq.items[priority]
		if len(items) > 0 {
			// if thrs anything in the queue
			// get the net item
			next := items[0]
			// reslice the priority queue
			pq.items[priority] = items[1:] // effectively deletes teh item at first index
			return next, true
		}
	}
	var empty V
	return empty, false // false means theres nth left in the priority queue
}

// return a priority queue that is generic over P and V
// this function is not a receiver fn hence we need to specifiy the generic constraints again
func NewPriorityQueue[P comparable, V any](priorities []P) PriorityQueue[P, V] {
	return PriorityQueue[P, V]{items: make(map[P][]V), priorities: priorities}
}

func main() {
	// integer as priority, string as items in the queue
	queue := NewPriorityQueue[int, string]([]int{High, Medium, Low})
	// we are creating a new integer slice and setting it to High Medium and Low

	queue.Add(Low, "L-1") // adding a low priority item
	queue.Add(High, "H-1")

	fmt.Println(queue.Next())

	queue.Add(Medium, "M-1")
	queue.Add(High, "H-2")
	queue.Add(High, "H-3")

	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())

}
