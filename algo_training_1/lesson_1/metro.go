package lesson_1

import "math"

func Metro(a, b, n, m int) (int, int) {
	min1 := a*(n-1) + n
	max1 := min1 + 2*a

	min2 := b*(m-1) + m
	max2 := min2 + 2*b

	min := int(math.Max(float64(min1), float64(min2)))
	max := int(math.Min(float64(max1), float64(max2)))

	if max < min {
		return -1, -1
	}

	return min, max
}
