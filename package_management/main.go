package main

import (
	"fmt"
	"package_management/test"
)

func main() {
	// Example of using the add function from the test package
	result, err := test.Add(5, 10)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result of addition:", result)
}
