package bucketservice

import (
	"math/rand"
	"time"
)

var seededRandom = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRandom.Intn(len(charset))]
	}
	return string(b)
}

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString Create random string.
func randomString(length int) string {
	return stringWithCharset(length, charset)
}
