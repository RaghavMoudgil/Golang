package main

import (
	"fmt"
	"time"
)

func task(ch chan<- string) {
	time.Sleep(3 * time.Second)
	ch <- "task done"
}

func main() {
	ch := make(chan string)
	go task(ch)
	select {
	case res := <-ch:
		fmt.Println("task complete", res)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: task is taking longer than expected")
	}
}
