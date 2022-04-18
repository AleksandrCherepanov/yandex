package lesson_1

func Gears(metal_mass, blank_mass, gear_mass int) int {
	if metal_mass < blank_mass {
		return 0
	}

	if blank_mass < gear_mass {
		return 0
	}

	if metal_mass < gear_mass {
		return 0
	}

	var result int

	for {
		metal_rest := 0
		blankCount := metal_mass / blank_mass
		if blankCount == 0 {
			break
		}

		metal_rest += metal_mass % blank_mass

		for i := 0; i < blankCount; i++ {
			result += blank_mass / gear_mass
			metal_rest += blank_mass % gear_mass
		}

		metal_mass = metal_rest
	}

	return result
}
