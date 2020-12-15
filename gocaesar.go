package gocaesar

import (
	"fmt"
	"github.com/zxxshaycormac/gocaesar/pkg/define"
	"strconv"
)

type Caesar struct {
	input      string
	numberStep int
	letterStep int
}

func NewCaesar() *Caesar {
	return new(Caesar)
}

func (this *Caesar) Input(input string) *Caesar {
	this.input = input
	return this
}

//设置数字步长
func (this *Caesar) NumberStep(step int) *Caesar {
	this.numberStep = step
	return this
}

//设置字母步长
func (this *Caesar) LetterStep(step int) *Caesar {
	this.letterStep = step
	return this
}

//同时设置字母和数字步长
func (this *Caesar) Step(step int) *Caesar {
	this.numberStep = step
	this.letterStep = step
	return this
}

//加密
func (this *Caesar) BeCaesar() string {
	var result = ""
	inputList := splitString(this.input)
	for _, v := range inputList {
		if num, err := strconv.Atoi(v); err == nil {
			newNumber := define.GetNumShift(num, this.numberStep)
			result = fmt.Sprintf("%s%v", result, newNumber)
		} else {
			newLetter := define.GetLetShift(v, this.letterStep)
			result = fmt.Sprintf("%s%s", result, newLetter)
		}
	}
	return result
}

//解密
func (this *Caesar) DeCaesar() string {
	var result = ""
	inputList := splitString(this.input)
	for _, v := range inputList {
		if num, err := strconv.Atoi(v); err == nil {
			newNumber := define.GetDeNumShift(num, this.numberStep)
			result = fmt.Sprintf("%s%v", result, newNumber)
		} else {
			newLetter := define.GetDeLetShift(v, this.letterStep)
			result = fmt.Sprintf("%s%s", result, newLetter)
		}
	}
	return result
}

func splitString(input string) []string {
	var list = make([]string, 0)
	for _, v := range []rune(input) {
		list = append(list, string(v))
	}
	return list
}
