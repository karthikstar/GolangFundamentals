package queue

// making a fifo queue
type Queue struct {
	items    []int
	capacity int
}

func New(capacity int) Queue {
	return Queue{make([]int, 0, capacity), capacity}
}

//append fn to add a item to the queue
// this is a receiver function
func (q *Queue) Append(item int) bool {
	if len(q.items) == q.capacity {
		return false // if queue is full we cant add items to it
	}
	q.items = append(q.items, item)
	return true
}

//retrieve next item from queue - returns next item
func (q *Queue) Next() (int, bool) {
	if len(q.items) > 0 {
		next := q.items[0]    // first one that goes in is what comes out first
		q.items = q.items[1:] //reslice
		return next, true
	} else {
		return 0, false // means queue is empty
	}
}
