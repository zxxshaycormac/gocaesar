# gocaesar
go实现的凯撒加解密

用法可以参考gocaesar_test.go

支持大小写字母和数字，支持字母数字位移不同位数，支持将密文夹在时间戳的base64字符中间进行混淆

加密

```go
text := "aBcD1234"
nStep := 9
lStep := 25
result := NewCaesar().Input(text).NumberStep(nStep).LetterStep(lStep).BeCaesar()
fmt.Printf("原文 : %s, 数字位移 : %d,字母位移 : %d\n", text, nStep, lStep)
fmt.Printf("密文 : %s", result)
//使用Base64()方法会自动取当前的nano时间戳，进行base64转成字符串后将密文拼接进你指定的下标位置
bit := 26
result = NewCaesar().Input(text).NumberStep(nStep).LetterStep(lStep).Base64(bit).BeCaesar()
fmt.Printf("原文 : %s, 数字位移 : %d,字母位移 : %d, base64填充下标位置 : %d", text, nStep, lStep, bit)
fmt.Printf("密文 : %s", result)
```

字母和数字位移位数相同

```go
text := "aBcD1234"
step := 9
result := NewCaesar().Input(text).Step(step).BeCaesar()
fmt.Printf("原文 : %s, 位移 : %d\n", text, step)
fmt.Printf("密文 : %s", result)
```

解密

```go
text := "zAbC0123"
nStep := 9
lStep := 25
result := NewCaesar().Input(text).NumberStep(nStep).LetterStep(lStep).DeCaesar()
fmt.Printf("密文 : %s, 数字位移 : %d,字母位移 : %d\n", text, nStep, lStep)
fmt.Printf("原文 : %s", result)
//使用Base64()方法会自动根据你输入的下标，把加密时拼接的混淆用base64字符剪掉
bit := 26
text = "MTYwODA5OTU5NDc3NzQxMzMwMAzAbC0123=="
result = NewCaesar().Input(text).NumberStep(nStep).LetterStep(lStep).Base64(bit).DeCaesar()
fmt.Printf("密文 : %s, 数字位移 : %d,字母位移 : %d, base64填充下标位置 : %d", text, nStep, lStep, bit)
fmt.Printf("原文 : %s", result)
```

