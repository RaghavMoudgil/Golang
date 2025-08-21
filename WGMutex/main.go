package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Starting main function")
	go func() {
		// time.Sleep(2 * time.Second)
		fmt.Println("Goroutine finished after 2 seconds")
		wg.Done()

	}()
	go func() {
		// time.Sleep(4 * time.Second)
		fmt.Println("Goroutine finished after 3 seconds")
		wg.Done()

	}()

	// bwlow is the code to use mutex
	//just comment thr ccde with the mutex to see the difference
	var mu sync.Mutex
	var count int
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			fmt.Println("this is the value of i", i)
			count++
			mu.Unlock()
		}()
	}
	fmt.Println("Count value:", count)
	// wg.Done()

	defer wg.Wait()

}
