package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Generating random number using Cryptography")

	random_num, err := rand.Int(rand.Reader, big.NewInt(54))
	if err != nil {
		panic(err)
	}
	fmt.Println(random_num)
}
