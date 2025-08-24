package main

import (
	"fmt"
	"sync"

	"golang.org/x/tools/go/analysis/passes/defers"
)

// in this multiple go routines are going to be send through a single channel to merge the results

func worker(ch chan<- string, msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch<-

}
func main() {
	ch := make(chan string)

	var wg *sync.WaitGroup

	message := []string{"this","is","fan","in","pattern"}

	wg.Add(len(message))
	//this is where we are allocating the worker
	for _, v := range message {
		go func ()  {
			worker(ch,v,wg)
		}()
	}

	// here we are going to close the ch

	go func ()  {
		wg.Wait()
		close(ch)

	}()

	//printing out the data
	for _, msg := range ch {
		fmt.Print(msg)
	}


}
