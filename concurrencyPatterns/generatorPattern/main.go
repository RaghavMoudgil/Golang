package main

import (
	"fmt"
)

func generator() <-chan int {

	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch

}

func main() {
	fmt.Println("This is the main package for the generator pattern.")

	ch := generator()

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

}

//generator func simple create channels and supplies them on demand one by one
