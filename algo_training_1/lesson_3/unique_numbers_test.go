package lesson_3

import "testing"

func TestChampionship(t *testing.T) {
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
			1,
		},
		{
			"case #3",
			[]int{1, 2},
			2,
		},
		{
			"case #4",
			[]int{1, 2, 3, 2, 1},
			3,
		},
		{
			"case #5",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			10,
		},
		{
			"case #6",
			[]int{1, 2, 3, 4, 5, 1, 2, 1, 2, 7, 3},
			6,
		},
	}

	for _, tc := range testCases {
		result := UniqueNumbers(tc.input)

		if result != tc.expected {
			t.Errorf("\n%v failed.\n Expected: %v.\n Actual: %v.", tc.name, tc.expected, result)
		}
	}
}
