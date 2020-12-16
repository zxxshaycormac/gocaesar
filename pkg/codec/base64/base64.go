package base64

import (
	"encoding/base64"
	"strings"
)

func Encode2Base64Str(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func DecodeBase64Str(base64Str string) ([]byte, error) {
	// 如果包含 `-`或者`_`，则认为使用的是 URLEncoding
	if strings.Contains(base64Str, "_") || strings.Contains(base64Str, "-") {
		return base64.URLEncoding.DecodeString(base64Str)
	}
	return base64.StdEncoding.DecodeString(base64Str)
}
