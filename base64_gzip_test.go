package goli_test

import (
	"testing"

	"github.com/oliverisaac/goli"
	"golang.org/x/exp/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestGzipAndBase64Encode(t *testing.T) {
	t.Run("test unzip and rezip", func(t *testing.T) {
		input := RandStringRunes(100)
		compressed, err := goli.GzipAndBase64Encode(input)
		if err != nil {
			t.Errorf("GzipAndBase64Encode() error = %v", err)
			return
		}
		decompressed, err := goli.Base64DecodeAndGunzip(compressed)
		if err != nil {
			t.Errorf("GzipAndBase64Encode() error = %v", err)
			return
		}

		if decompressed != input {
			t.Fatalf("Compressed and uncompressed did not match: input: %q, decompressed: %q", input, decompressed)
		}
	})
}
