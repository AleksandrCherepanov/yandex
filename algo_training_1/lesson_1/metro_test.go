package lesson_1

import "testing"

func TestMetro(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"case #1",
			[]int{1, 3, 3, 2},
			[]int{5, 7},
		},
		{
			"case #2",
			[]int{1, 5, 1, 2},
			[]int{-1, -1},
		},
	}

	for _, tc := range testCases {
		r1, r2 := Metro(tc.input[0], tc.input[1], tc.input[2], tc.input[3])

		if r1 != tc.expected[0] {
			t.Errorf("\n%v failed.\n Expected: %v.\n Actual: %v.", tc.name, tc.expected, r1)
		}

		if r2 != tc.expected[1] {
			t.Errorf("\n%v failed.\n Expected: %v.\n Actual: %v.", tc.name, tc.expected, r2)
		}
	}
}
