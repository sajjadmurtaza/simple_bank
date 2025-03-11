package util

import (
	"fmt"
	"math/rand"
	"strings"
)

func init() {
	fmt.Println("rand.Seed() deprecated...")
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(n int) string {
	var stringBuilder strings.Builder

	charsetLength := len(charset)

	for i := 0; i < n; i++ {
		randomPositionChar := charset[rand.Intn(charsetLength)]
		stringBuilder.WriteByte(randomPositionChar)
	}

	return stringBuilder.String()
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "JPY"}

	return currencies[rand.Intn(len(currencies))]
}
