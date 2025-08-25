package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	sem <- struct{}{}

	fmt.Println("work started", id)
	time.Sleep(time.Second)
	fmt.Println("work ended", id)
	<-sem

}
func main() {
	const limit = 2
	sem := make(chan struct{}, limit)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go task(i, sem, &wg)
	}
	wg.Wait()
}
