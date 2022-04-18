package lesson_1

func sliceMax(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	max := slice[0]
	for _, num := range slice {
		if num > max {
			max = num
		}
	}

	return max
}

func Laptops(a, b, c, d int) (int, int) {
	options := make([][]int, 0)

	options = append(options, []int{a + c, sliceMax([]int{b, d})})
	options = append(options, []int{a + d, sliceMax([]int{b, c})})
	options = append(options, []int{b + c, sliceMax([]int{a, d})})
	options = append(options, []int{b + d, sliceMax([]int{a, c})})

	intSize := 32 << (^uint(0) >> 63)
	min := 1<<(intSize-1) - 1
	var ts1, ts2 int

	for _, option := range options {
		area := option[0] * option[1]
		if area < min {
			min = area
			ts1 = option[0]
			ts2 = option[1]
		}
	}

	return ts1, ts2
}
