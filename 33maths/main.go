package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("We are going to learn about Maths, Crypto and random functions in Golang")
	//for creatign a random no. ew will need to use rand function from maths
	//we need to seed this rand func ro the algorithm run by this random function
	//if we keep the seed same it will alaways give me the samw number so we need to give a number that is changing all the time foreg. TIME
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(54))

}
