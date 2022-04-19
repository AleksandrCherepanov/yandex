package lesson_1

import "testing"

func TestCastle(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected string
	}{
		{
			"case #1",
			[]int{1, 1, 1, 1, 1},
			"YES",
		},
		{
			"case #2",
			[]int{2, 2, 2, 1, 1},
			"NO",
		},
	}

	for _, tc := range testCases {
		result := Castle(tc.input[0], tc.input[1], tc.input[2], tc.input[3], tc.input[4])

		if result != tc.expected {
			t.Errorf("\n%v failed.\n Expected: %v.\n Actual: %v.", tc.name, tc.expected, result)
		}
	}
}
