package lesson_3

func UniqueNumbers(input []int) int {
	unique := make(map[int]bool, 0)
	count := 0

	for _, num := range input {
		if _, ok := unique[num]; !ok {
			unique[num] = true
			count++
		}
	}

	return count
}
