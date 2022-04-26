package lesson_2

func Neighbours(input []int) int {
	result := 0

	for i := 0; i < len(input); i++ {
		left := i - 1
		right := i + 1

		if left < 0 {
			left = i
		}

		if right >= len(input) {
			right = i
		}

		if input[i] > input[left] && input[i] > input[right] {
			result++
		}
	}

	return result
}
