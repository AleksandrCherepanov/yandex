package lesson_1

import "strings"

func Phones(phones []string) []string {
	result := make([]string, 0, len(phones))

	for i := 0; i < len(phones); i++ {
		phones[i] = strings.ReplaceAll(phones[i], "-", "")
		phones[i] = strings.ReplaceAll(phones[i], "(", "")
		phones[i] = strings.ReplaceAll(phones[i], ")", "")
		phones[i] = strings.ReplaceAll(phones[i], "+7", "8")

		if len(phones[i]) == 7 {
			phones[i] = "495" + phones[i]
		}
		if len(phones[i]) == 10 {
			phones[i] = "8" + phones[i]
		}

		if i > 0 {
			if phones[i] == phones[0] {
				result = append(result, "YES")
			} else {
				result = append(result, "NO")
			}
		}
	}

	return result
}
