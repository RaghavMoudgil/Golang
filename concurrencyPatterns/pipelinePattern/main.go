package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is the code for pipeline")

	ch := make(chan int)
	result := calSquare(ch)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range result {
		fmt.Println(v)
	}

}

func calSquare(ch chan int) chan int {
	result := make(chan int)
	go func() {
		for v := range ch {
			result <- v * v
		}
		close(result)
	}()
	return result
}
