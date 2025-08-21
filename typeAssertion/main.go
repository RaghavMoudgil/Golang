package main

import (
	"fmt"
)

func main() {

	// below is the example of type assertion
	var i interface{} = "hello"

	a := i.(string)
	fmt.Println(a)

	//nelow is the example of type switching
	switch v := i.(type) {
	case string:
		fmt.Println("String:", v)
	case int:
		fmt.Println("Integer:", v)
	default:
		fmt.Println("Unknown type")
	}
}
