package lesson_1

func Castle(a, b, c, d, e int) string {
	switch true {
	case a <= d && c <= e:
		return "YES"
	case a <= d && b <= e:
		return "YES"
	case b <= d && c <= e:
		return "YES"
	case c <= d && a <= e:
		return "YES"
	case b <= d && a <= e:
		return "YES"
	case c <= d && b <= e:
		return "YES"
	}

	return "NO"
}
