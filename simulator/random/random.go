package random

import (
	"math/rand"
	"time"
)

type GenerateRandomIntFunction func(int, int) int

func GenerateRandomInt(min int, max int) int {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min) + min
}
