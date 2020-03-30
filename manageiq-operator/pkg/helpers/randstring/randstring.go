package randstring

import (
	"math/rand"
	"time"
)

const EncryptionCharSet = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const GUIDCharSet = "abcdefghijklmnopqrstuvwxyz0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateEncryptionKey(length int) string {
	return StringWithCharset(length, EncryptionCharSet)
}

func GenerateGUID() string {
	return StringWithCharset(8, GUIDCharSet) + "-" + StringWithCharset(4, GUIDCharSet) + "-" + StringWithCharset(4, GUIDCharSet) + "-" + StringWithCharset(4, GUIDCharSet) + "-" + StringWithCharset(12, GUIDCharSet)
}