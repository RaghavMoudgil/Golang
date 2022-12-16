package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Thsi is the test code for general purpose \n")
	//First we are foinf to check about the Exit statemenr
	//and difference between exit and defer statement
	defer fmt.Println("Defer is exicuted !!!")
	os.Exit(4) //beacuse of this exit statement defer is not exicuted here !!!
}
