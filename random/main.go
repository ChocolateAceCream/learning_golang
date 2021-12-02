package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

func randomID() string {
	return uuid.New().String()
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func main() {
	uuid := randomID()
	fmt.Println(uuid)
	rand64 := randomFloat(1, 10)
	fmt.Println(rand64)
}
