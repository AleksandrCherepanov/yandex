package lesson_2

import "testing"

func TestNeighboors(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"case #1",
			[]int{},
			0,
		},
		{
			"case #2",
			[]int{1},
			0,
		},
		{
			"case #3",
			[]int{1, 2},
			0,
		},
		{
			"case #4",
			[]int{1, 2, 3},
			0,
		},
		{
			"case #5",
			[]int{1, 2, 1},
			1,
		},
		{
			"case #6",
			[]int{1, 2, 3, 4, 5},
			0,
		},
		{
			"case #7",
			[]int{5, 4, 3, 2, 1},
			0,
		},
		{
			"case #8",
			[]int{1, 5, 1, 5, 1},
			2,
		},
	}

	for _, tc := range testCases {
		result := Neighbours(tc.input)

		if result != tc.expected {
			t.Errorf("\n%v failed.\n Expected: %v.\n Actual: %v.", tc.name, tc.expected, result)
		}
	}
}
