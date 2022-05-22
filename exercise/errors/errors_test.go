package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	// use a test table so taht we can test multiple inouts
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"1:3:44", true},
		{"bad", false},
		{"1:-3:44", false},
		{"0:59:59", true},
		{"", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			// means if data ok is true but theres a error
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}
