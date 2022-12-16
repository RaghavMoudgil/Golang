package main

import "fmt"

func val() (int, int) {
	return 6, 5
}
func main() {
	fmt.Println("Returning multiple values from the function")
	a, b := val()
	fmt.Println(a)
	fmt.Println(b)
	c, _ := val() //if you leave fist underscore it will print the value of second and vise-versa
	fmt.Println(c)
}
