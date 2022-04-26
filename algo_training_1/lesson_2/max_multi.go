package lesson_2

func MaxMulti(input []int) []int {
	max1 := 0
	max2 := 0

	min1 := 0
	min2 := 0

	if input[0] > input[1] {
		max1 = input[0]
		max2 = input[1]

		min1 = input[1]
		min2 = input[0]
	} else {
		max1 = input[1]
		max2 = input[0]

		min1 = input[0]
		min2 = input[1]
	}

	for i := 2; i < len(input); i++ {
		if input[i] > max1 {
			max2 = max1
			max1 = input[i]
		} else if input[i] > max2 {
			max2 = input[i]
		}
	}

	for i := 2; i < len(input); i++ {
		if input[i] < min1 {
			min2 = min1
			min1 = input[i]
		} else if input[i] < min2 {
			min2 = input[i]
		}
	}

	minMulti := min1 * min2
	maxMulti := max1 * max2

	if minMulti > maxMulti {
		return []int{min1, min2}
	}

	return []int{max2, max1}
}
