package helper

import (
	"log"
	"math/rand"
	"time"
)

func LogString(s string) {
	log.Println("[LOG]", s)
}

func LogInt(n int) {
	log.Println("[LOG]", n)
}

func GenerateRandom(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}
