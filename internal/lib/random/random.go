package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {
	charset := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, size)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

// TODO: Make good randomizer
