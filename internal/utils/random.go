package utils

import (
	"math/rand"
	"strings"
	"time"
)

var r *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomCurrency() string {
	currencies := []string{"IDR", "MYR", "SGD", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
