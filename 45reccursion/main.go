package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) //fact is a self calling fucntion
}

func main() {
	fmt.Println("Code to learn about reccursion.")
	fmt.Println(fact(6))

}
