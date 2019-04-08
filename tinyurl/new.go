package tinyurl

import (
	"math/rand"
	"time"
)

func CreateRandomString() string {
	const urlLength int = 8
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	rand.Seed(time.Now().UnixNano())

	r := make([]rune, urlLength)
	for i := range r {
		r[i] = letters[rand.Intn(len(letters))]
	}
	return string(r)
}
