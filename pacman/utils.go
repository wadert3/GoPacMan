package pacman

func getMirrorLocation(value int, length int) int {
	length--
	if value > length {
		return length - (value - length - 1)
	}

	return value
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isValidIndex(valueArray []int) bool {
	for _, value := range valueArray {
		if value < 0 {
			return false
		}
	}
	return true
}
