package gocaesar

import (
	"fmt"
	"shay/gocaesar/pkg/define"
	"strconv"
)

//加密
func BeCaesar(input string, nStep, lStep int) string {
	var result string = ""
	inputList := SplitString(input)
	for _, v := range inputList {
		if num, err := strconv.Atoi(v); err == nil {
			newNumber := define.GetNumShift(num, nStep)
			result = fmt.Sprintf("%s%v", result, newNumber)
		} else {
			newLetter := define.GetLetShift(v, lStep)
			result = fmt.Sprintf("%s%s", result, newLetter)
		}
	}
	return result
}

//解密
func DeCaesar(input string, nStep, lStep int) string {
	var result string = ""
	inputList := SplitString(input)
	for _, v := range inputList {
		if num, err := strconv.Atoi(v); err == nil {
			newNumber := define.GetDeNumShift(num, nStep)
			result = fmt.Sprintf("%s%v", result, newNumber)
		} else {
			newLetter := define.GetDeLetShift(v, lStep)
			result = fmt.Sprintf("%s%s", result, newLetter)
		}
	}
	return result
}

func SplitString(input string) []string {
	var list = make([]string, 0)
	for _, v := range []rune(input) {
		list = append(list, string(v))
	}
	return list
}
