package main

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		max := big.NewInt(int64(len(charset)))
		randIndex, _ := rand.Int(rand.Reader, max)
		b[i] = charset[randIndex.Int64()]
	}
	return string(b)
}
