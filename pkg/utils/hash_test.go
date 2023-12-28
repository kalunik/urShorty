package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	const (
		hashMaxLen   = 8
		allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	)
	hash, err := GenerateHash()

	assert.Nil(t, err, "error wasn't expected")
	assert.NotEmptyf(t, hash, "hash so can't be empty : got '%s'", hash)
	assert.Lenf(t, hash, hashMaxLen, "hash's lenght should be %d", hashMaxLen)
	for _, v := range hash {
		assert.Contains(t, allowedChars, string(v), "hash shouldn't contain not allowed char '%c'", v)
	}
}
