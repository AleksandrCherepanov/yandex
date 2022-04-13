package lesson_1

func AirConditioner(troom, tcond int, mode string) int {
	if mode == "fan" {
		return troom
	}

	if mode == "auto" {
		return tcond
	}

	if mode == "heat" {
		if tcond > troom {
			return tcond
		}

		return troom
	}

	if mode == "freeze" {
		if tcond < troom {
			return tcond
		}
	}

	return troom
}
