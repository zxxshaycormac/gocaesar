package gocaesar

import "testing"

func TestBeCaesar(t *testing.T) {
	text := "abcd1234"
	nStep := 9
	lStep := 25
	result := BeCaesar(text, nStep, lStep)
	t.Logf("原文 : %s, 数字位移 : %d,字母位移 : %d", text, nStep, lStep)
	t.Logf("密文 : %s", result)
}

func TestDeCaesar(t *testing.T) {
	text := "zabc0123"
	nStep := 9
	lStep := 25
	result := DeCaesar(text, nStep, lStep)
	t.Logf("密文 : %s, 数字位移 : %d,字母位移 : %d", text, nStep, lStep)
	t.Logf("原文 : %s", result)
}