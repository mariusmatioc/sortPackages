package pkg

import "testing"

func TestSort(t *testing.T) {
	var tests = []struct {
		width, height, length, mass float64
		want                        string
	}{
		{0, 10, 30, -1, REJECTED},
		{151, 30, 20, 10, SPECIAL},
		{10, 30, 20, 10, STANDARD},
		{1000, 1000, 1, 10, SPECIAL},
		{1000, 1000, 1, 25, REJECTED},
		{10, 10, 25, 25, SPECIAL},
	}
	for _, test := range tests {
		if got := Sort(test.width, test.height, test.length, test.mass); got != test.want {
			t.Errorf("sort(%f,%f,%f,%f) = %s", test.width, test.height, test.length, test.mass, got)
		}
	}
}
