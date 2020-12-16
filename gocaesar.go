package gocaesar

import (
	"fmt"
	"github.com/zxxshaycormac/gocaesar/pkg/define"
	"github.com/zxxshaycormac/gocaesar/pkg/codec/base64"
	"strconv"
	"time"
)

type Caesar struct {
	input      string
	numberStep int
	letterStep int
	hash       bool
	hashType   string
	hashBit    int
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

func (this *Caesar) Base64(bit int) *Caesar {
	this.hash = true
	this.hashType = "base64"
	if bit > 26 {
		bit = 26
	} else if bit < 0 {
		newBit := bit % 26
		bit = 26 - newBit
	}
	this.hashBit = bit
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
	if this.hash {
		switch this.hashType {
		case "base64":
			return this.hashBase64(result)
		}
	}
	return result
}

//解密
func (this *Caesar) DeCaesar() string {
	var result = ""
	if this.hash {
		this.input = this.deHashBase64()
	}
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

func (this *Caesar) hashBase64(input string) string {
	theBase64 := nowBase64()
	return fmt.Sprintf("%s%s%s", theBase64[:this.hashBit], input, theBase64[this.hashBit:])
}

func (this *Caesar) deHashBase64() string {
	return this.input[this.hashBit : this.hashBit+(len(this.input)-28)]
}

func nowBase64() string {
	return base64.Encode2Base64Str([]byte(fmt.Sprintf("%v", time.Now().UnixNano())))
}
