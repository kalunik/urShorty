package utils

import (
	"bytes"
	"math/rand"
)

func GenerateHash(full string) (string, error) {

	const (
		base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		base        = 62
		shortLength = 8
	)

	shortURL := &bytes.Buffer{}

	for i := 0; i < shortLength; i++ {
		err := shortURL.WriteByte(base62Chars[rand.Intn(base)])
		if err != nil {
			return "", err
		}
	}

	return shortURL.String(), nil
}
