package main

import (
	"fmt"
	"sync"
)

// in this multiple go routines are going to be send through a single channel to merge the results

func worker(ch chan<- string, msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- msg
}
func main() {
	ch := make(chan string)

	var wg sync.WaitGroup

	message := []string{"this", "is", "fan", "in", "pattern"}

	wg.Add(len(message))
	//this is where we are allocating the worker
	for _, v := range message {
		go worker(ch, v, &wg)
	}

	// here we are going to close the ch

	go func() {
		wg.Wait()
		close(ch)

	}()

	//printing out the data
	for msg := range ch {
		fmt.Print(msg, " ")
	}

}
