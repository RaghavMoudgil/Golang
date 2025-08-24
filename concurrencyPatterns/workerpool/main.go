package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, job <-chan int, wg *sync.WaitGroup) {

	for v := range job {
		fmt.Printf("this is my work:Worker %d starting;job %d\n", id, v)
		time.Sleep(time.Second)

	}

	defer wg.Done()

}

func main() {
	fmt.Println("This is the main package for the worker pool pattern.")

	job := make(chan int, 5)

	var wg sync.WaitGroup

	//here i am starting 4 workers
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(i, job, &wg)

	}

	//here i am sending josbs to cahnnel 5
	for i := 0; i < 5; i++ {
		job <- i

	}

	close(job)
	wg.Wait()

}
