package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var currencies, c = []string{
	USD,
	EUR,
	JPY,
	GBP,
	CNY,
	AUD,
	CAD,
	CHF,
	HKD,
	RUB,
}, len(currencies)

// init sets new source for pseudo-random
func init() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

// RandomInt64 generates a random 64-bit integer between min and max
func RandomInt64(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// randomElement returns random element from slice
func randomElement(container []string, len int) string {
	return container[rand.Intn(len)]
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt64(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	return randomElement(currencies, c)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
