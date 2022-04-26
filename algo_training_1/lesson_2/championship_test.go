package lesson_2

import "testing"

func TestChampionship(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"case #1",
			[]int{15, 3, 7},
			0,
		},
		{
			"case #2",
			[]int{1, 2, 3},
			0,
		},
		{
			"case #3",
			[]int{10, 20, 15, 10, 30, 5, 1},
			6,
		},
		{
			"case #4",
			[]int{15, 15, 10},
			1,
		},
		{
			"case #5",
			[]int{10, 15, 20},
			0,
		},
	}

	for _, tc := range testCases {
		result := Championship(tc.input)

		if result != tc.expected {
			t.Errorf("\n%v failed.\n Expected: %v.\n Actual: %v.", tc.name, tc.expected, result)
		}
	}
}
