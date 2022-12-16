package main

import "fmt"

func val(pval int) {
	pval = 0
}
func pointer(pptr *int) {
	*pptr = 0
}
func main() {
	fmt.Println("Code to learn about the pointer.")
	i := 5
	fmt.Println("intial value is :", i)
	val(i)
	fmt.Println("value of the pointer :", i)
	pointer(&i)
	fmt.Println("pointer value(pval) :", i)
	fmt.Println("pointer:", &i)
}
