package goli

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// GenerateRandomString generates a random string of length = n
// You may pass in one or more strings to use as the default character set to select from
// The default character set is a restricted set of letters+numbers to help avoid common bad words
//
//	GenerateRandomString(10) // Generate string with length of 10
//	GenerateRandomString(10, "abc1123") // Generate string with only letters abc123
func GenerateRandomString(n int, charsets ...string) (string, error) {
	letters := "2346789ABCDEGHJKLMPQRTVWXYZabcdeghijkmpqrtuvwxyz"
	if len(charsets) > 0 {
		letters = strings.Join(charsets, "")
	}
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
