package main

import "fmt"


func main() {
	fmt.Println("This is the code for structs in GO lang")
	Raghav := User{"Raghav",15,true}
	fmt.Printf("Raghav's Details: %+v\n" ,Raghav)
	Raghav.GetStatus()
	Raghav.Getage()
}
type User struct{
	Name string 
	Age int 
	Status bool
}
func (u User) GetStatus()  {
	fmt.Println("is User active:" ,u.Status)
}
func(u User) Getage(){
	fmt.Println("user age is: ", u.Age)
}