package utils

import (
	"math/rand"
	"time"
)

func RandomName() string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbn")
	name := make([]byte, 10)
	rand.Seed(time.Now().Unix())
	for i := range name {
		name[i] = letters[rand.Intn(10)]
	}
	return string(name)
}
