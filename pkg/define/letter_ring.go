package define

var lring = map[int]string{
	0:  "a",
	1:  "b",
	2:  "c",
	3:  "d",
	4:  "e",
	5:  "f",
	6:  "g",
	7:  "h",
	8:  "i",
	9:  "j",
	10: "k",
	11: "l",
	12: "m",
	13: "n",
	14: "o",
	15: "p",
	16: "q",
	17: "r",
	18: "s",
	19: "t",
	20: "u",
	21: "v",
	22: "w",
	23: "x",
	24: "y",
	25: "z",
}

var reverseLring = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
	"i": 8,
	"j": 9,
	"k": 10,
	"l": 11,
	"m": 12,
	"n": 13,
	"o": 14,
	"p": 15,
	"q": 16,
	"r": 17,
	"s": 18,
	"t": 19,
	"u": 20,
	"v": 21,
	"w": 22,
	"x": 23,
	"y": 24,
	"z": 25,
}

var capitalLring = map[int]string{
	0:  "A",
	1:  "B",
	2:  "C",
	3:  "D",
	4:  "E",
	5:  "F",
	6:  "G",
	7:  "H",
	8:  "I",
	9:  "J",
	10: "K",
	11: "L",
	12: "M",
	13: "N",
	14: "O",
	15: "P",
	16: "Q",
	17: "R",
	18: "S",
	19: "T",
	20: "U",
	21: "V",
	22: "W",
	23: "X",
	24: "Y",
	25: "Z",
}

var capitalReverseLring = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
	"I": 8,
	"J": 9,
	"K": 10,
	"L": 11,
	"M": 12,
	"N": 13,
	"O": 14,
	"P": 15,
	"Q": 16,
	"R": 17,
	"S": 18,
	"T": 19,
	"U": 20,
	"V": 21,
	"W": 22,
	"X": 23,
	"Y": 24,
	"Z": 25,
}

func GetLetShift(let string, step int) string {
	if _, ok := reverseLring[let]; ok {
		//小写字母
		return makeShift(reverseLring, lring, let, step)
	} else {
		//大写字母
		return makeShift(capitalReverseLring, capitalLring, let, step)
	}
}

func makeShift (reverseRing map[string]int, ring map[int]string, let string, step int) string{
	newLet := reverseRing[let] + step
	if newLet > 25 {
		return ring[newLet%26]
	}
	return ring[newLet]
}

func GetDeLetShift(let string, step int) string {
	if _, ok := reverseLring[let]; ok {
		//小写字母
		return makeDeShift(reverseLring, lring, let, step)
	} else {
		//大写字母
		return makeDeShift(capitalReverseLring, capitalLring, let, step)
	}
}

func makeDeShift(reverseRing map[string]int, ring map[int]string, let string, step int) string {
	newLet := reverseRing[let] - step
	if newLet < 0 {
		newLetUnit := newLet % 26
		if newLetUnit == 0 {
			return ring[newLetUnit]
		}
		return ring[26+newLetUnit]
	}
	return ring[newLet]
}
