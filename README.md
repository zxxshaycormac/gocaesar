# gocaesar
go实现的凯撒加解密

用法可以参考gocaesar_test.go

目前只支持小写字母和数字，支持字母数字位移不同位数

加密

```
text := "abcd1234"
nStep := 9
lStep := 25
result := NewCaesar().Input(text).NumberStep(nStep).LetterStep(lStep).BeCaesar()
fmt.Printf("原文 : %s, 数字位移 : %d,字母位移 : %d\n", text, nStep, lStep)
fmt.Printf("密文 : %s", result)
```

解密

```
text := "zabc0123"
nStep := 9
lStep := 25
result := NewCaesar().Input(text).NumberStep(nStep).LetterStep(lStep).DeCaesar()
fmt.Printf("密文 : %s, 数字位移 : %d,字母位移 : %d\n", text, nStep, lStep)
fmt.Printf("原文 : %s", result)
```

