package utils

import "math/rand"

func GenerateRandmoString(n int) string {
	var charsets = []rune("abcDEF1234567890abcdef")
	letters := make([]rune, n)
	for i := range letters {
		letters[i] = charsets[rand.Intn(len(charsets))]
	}
	return string(letters)
}
