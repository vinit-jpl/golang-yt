package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	// "time"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in GoLang")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.56

	// fmt.Println("The sum is:", myNumberOne+int(myNumberTwo))

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// random from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
