package lesson_2

func MaxIndex(slice []int) int {
	max := slice[len(slice)-1]
	maxIndex := len(slice) - 1
	for i := len(slice) - 2; i >= 0; i-- {
		if slice[i] >= max {
			maxIndex = i
			max = slice[i]
		}
	}

	return maxIndex
}

func Championship(input []int) int {
	result := 0
	found := false
	place := 0

	maxIndex := MaxIndex(input)

	for i := len(input) - 2; i > maxIndex; i-- {
		if found {
			break
		}
		if input[i]%10 == 5 && input[i+1] < input[i] {
			result++
			place = i
			found = true
		}
	}

	if !found {
		return result
	}

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] > input[place] {
			result++
		}
	}

	return result
}
