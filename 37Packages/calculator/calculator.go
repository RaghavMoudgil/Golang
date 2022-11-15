package calculator

import (
	"fmt"
)

func Cal() {
	fmt.Println("Enter the no. against the operation you want to perform.")
	fmt.Printf(" 1. Add \n 2. Subtract \n 3. Multiply \n 4. Divide \n")
	var choice int
	fmt.Scan(&choice)
	if choice == 1 || choice == 2 || choice == 3 || choice == 4 {

		var a float64
		fmt.Println("Enter first Number:")
		fmt.Scan(&a)
		var b float64
		fmt.Println("Enter second Number:")
		fmt.Scan(&b)
		var result float64
		if choice == 1 {
			result = a + b
		} else if choice == 2 {
			result = a - b
		} else if choice == 3 {
			result = a * b
		} else if choice == 4 {
			result = a / b
		}
		fmt.Println("Your result is:", result)
	} else {
		fmt.Println("Invalid Input")
	}
}
