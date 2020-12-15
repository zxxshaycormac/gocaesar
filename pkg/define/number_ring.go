package define

func GetNumShift(num, step int) int {
	newNum := num + step
	if newNum > 9 {
		return newNum % 10
	}
	return newNum
}

func GetDeNumShift(num, step int) int {
	newNum := num - step
	if newNum < 0 {
		newNumUnit := newNum % 10
		if newNumUnit == 0 {
			return 0
		}
		return 10 + newNumUnit
	}
	return newNum
}
