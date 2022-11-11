package main

import "fmt"

func main() {
	fmt.Println("Code to understand Defer.")
	//defer is the keyword which will help the line to execute at the end of the function
	//If there are multiple defer they will execute in LIFO format
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	newdefer()
}
func newdefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}