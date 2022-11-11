package main

import "fmt"

func main() {
	fmt.Println("Code for Pointers")
	// var ptr *int 
	// fmt.Println("Default Value of pointer is ", ptr)

	someNumber := 35
	var ptr = &someNumber // Defining value to the pointer
	fmt.Println("Value of pointer is :",ptr) // This is the memory adress
	fmt.Println("Value of pointer is :",*ptr) // This is the value inside the pointer

	*ptr = *ptr * 2
	fmt.Println("new value is ", someNumber)
	
}