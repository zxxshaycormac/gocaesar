package gocaesar

import (
	"fmt"
	"github.com/zxxshaycormac/gocaesar/pkg/codec/base64"
	"testing"
	"time"
)

func TestBeCaesar(t *testing.T) {
	text := "aBcD1234"
	nStep := 9
	lStep := 25
	result := NewCaesar().Input(text).NumberStep(nStep).LetterStep(lStep).BeCaesar()
	t.Logf("原文 : %s, 数字位移 : %d,字母位移 : %d", text, nStep, lStep)
	t.Logf("密文 : %s", result)
	bit := 26
	result = NewCaesar().Input(text).NumberStep(nStep).LetterStep(lStep).Base64(bit).BeCaesar()
	t.Logf("原文 : %s, 数字位移 : %d,字母位移 : %d, base64填充下标位置 : %d", text, nStep, lStep, bit)
	t.Logf("密文 : %s", result)
}

func TestDeCaesar(t *testing.T) {
	text := "zAbC0123"
	nStep := 9
	lStep := 25
	result := NewCaesar().Input(text).NumberStep(nStep).LetterStep(lStep).DeCaesar()
	t.Logf("密文 : %s, 数字位移 : %d,字母位移 : %d", text, nStep, lStep)
	t.Logf("原文 : %s", result)
	bit := 26
	text = "MTYwODA5OTU5NDc3NzQxMzMwMAzAbC0123=="
	result = NewCaesar().Input(text).NumberStep(nStep).LetterStep(lStep).Base64(bit).DeCaesar()
	t.Logf("密文 : %s, 数字位移 : %d,字母位移 : %d, base64填充下标位置 : %d", text, nStep, lStep, bit)
	t.Logf("原文 : %s", result)
}

func TestNowBase64(t *testing.T) {
	t.Log(time.Now().UnixNano())
	t.Log(base64.Encode2Base64Str([]byte(fmt.Sprintf("%v", time.Now().UnixNano()))))
}
