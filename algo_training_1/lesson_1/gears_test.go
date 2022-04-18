package lesson_1

import "testing"

func TestGears(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"case #1",
			[]int{10, 5, 2},
			4,
		},
		{
			"case #2",
			[]int{13, 5, 3},
			3,
		},
		{
			"case #3",
			[]int{200, 5, 200},
			4,
		},
	}

	for _, tc := range testCases {
		result := Gears(tc.input[0], tc.input[1], tc.input[2])

		if result != tc.expected {
			t.Errorf("\n%v failed.\n Expected: %v.\n Actual: %v.", tc.name, tc.expected, result)
		}
	}
}
