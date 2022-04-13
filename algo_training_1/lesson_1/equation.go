package lesson_1

import (
	"math"
	"strconv"
)

func Equation(a, b, c int) string {
	if a == 0 {
		t := int(math.Sqrt(float64(b)))
		if t != c {
			return "NO SOLUTION"
		}
		return "MANY SOLUTIONS"
	}

	if c < 0 {
		return "NO SOLUTION"
	}

	d := c*c - b
	x := d / a
	if x*a+b < 0 {
		return "NO SOLUTION"
	}

	if int(math.Sqrt(float64(a*x+b))) != c {
		return "NO SOLUTION"
	}

	return strconv.Itoa(x)
}
