package util

import (
	"math/rand/v2"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for range n {
		c := alphabet[rand.IntN(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int64N(max-min+1)
}