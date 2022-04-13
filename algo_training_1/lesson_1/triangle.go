package lesson_1

func Triangle(a, b, c int) string {
	if a < 0 || b < 0 || c < 0 {
		return "NO"
	}

	if a < b+c && b < a+c && c < a+b {
		return "YES"
	}

	return "NO"
}
