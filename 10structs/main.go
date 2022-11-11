package main

import "fmt"

func main() {
	fmt.Println("This is the code for structs in GO lang")
	// No classses are used in the GOlang instead struct is used 
	//no inheritance in go lang 
	Raghav := User{"Raghav",15,true}
	fmt.Printf("Raghav's Details: %+v\n" ,Raghav)//user +v for structure it will give more details
}
type User struct{
	Name string 
	Age int 
	Status bool
}