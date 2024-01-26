package util

import (
	"math/rand"
	"strings"
)

func RandStr(length int) string {
	const n = 62 // 26+26+10 (a-zA-Z0-9)
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteRune(nChar(rand.Int31n(n)))
	}
	return sb.String()
}

func nChar(n int32) rune {
	// 0-25
	// 26-51
	// 52-62
	if n < 26 {
		return n + 'a'
	}
	if n < 52 {
		return n - 26 + 'A'
	}
	return n - 52 + '0'
}

func RandBytes(length int) []byte {
	var result = make([]byte, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, byte(rand.Intn(256)))
	}
	return result
}
