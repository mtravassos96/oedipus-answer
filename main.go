package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("First tests")

	// Random time method
	randomTime := rand.New(rand.NewSource(time.Now().UnixNano())) 
	randomInt := randomTime.Intn(100)
	fmt.Print(randomInt)
}