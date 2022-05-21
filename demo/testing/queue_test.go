package queue

import "testing"

// test function that adds item to queue

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element %v, want %v", len(q.items), i)
		}
		if !q.Append(i) { // append will return true if succesful, and false if unsuccesful
			t.Errorf("failed to append item %v to queue", i) // this shld work for i = 0 to i = 2 as only 3 times
		}
	}
	if q.Append(4) { // if can append a 4th item, sths wrong w our queue as we are supp to only be able to add 3 things
		t.Errorf("should not be able to add to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}
	// 3 items in queue at this pt
	// 0 1 2 added into queue
	// 0 shld come out first

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Should be able to get item from queue")
		}
		if item != i { // item is supposed to equal to i. as first NexT() will give 0, then 1 then 2
			t.Errorf("got item in wrong order: %v, want %v", item, i)
		}
	}

	// Queue is empty at this point - so if theres any item left, sth wrong w the function
	item, ok := q.Next()
	if ok {
		t.Errorf("should not be any more items in queue, got : %v", item)
	}

}
