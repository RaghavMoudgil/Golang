package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("We are going to understand about the Race Condition in Golang")

	wg := &sync.WaitGroup{}
	mut := sync.Mutex{} //declairing mutex to get over the race consition
	var score = []int{0}
	//Declairing IFIes (imigiatly invoked functions just as in JavaScript )
	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, &mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, &mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, &mut)
	wg.Wait()
	fmt.Println(score)
}
