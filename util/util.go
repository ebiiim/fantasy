package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var randStrLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStr(n int) string {
	p := make([]rune, n)
	for idx := range p {
		p[idx] = randStrLetters[rand.Intn(len(randStrLetters))]
	}
	return string(p)
}

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
