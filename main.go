package main

import (
	"fmt"
	"math/big"
	//"math/rand"
	"crypto/rand"
	//"time"
)

func main() {
	fmt.Println("First tests")

	// Random time method
	// randomTime := rand.New(rand.NewSource(time.Now().UnixNano())) 
	// randomInt := randomTime.Intn(100)
	// fmt.Print(randomInt)

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}