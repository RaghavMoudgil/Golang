package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This code will help us to understand the Channels in Golang")
	//making channel
	mych := make(chan int, 2) //this 2 is the buffer values we can use in the channel simply means we can use max 2 value in the channel if we declair more we can save more values in a single channel
	wg := &sync.WaitGroup{}
	//mych <- 5
	//fmt.Println(<-mych)  //this will create a deadlock!

	wg.Add(2)
	//by using <- in the fuction it will make the channel SEND ONLY otherwise channel is bidirectional
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-mych

		fmt.Println("Is your channel open:", isChannelOpen)//we are usign this here because if the channel is close it will give us 0 value 
		fmt.Println("The value of channel is:", val)//and we will not be able to recognize weather it is a value that is passed or it is the channel value so to remove that confusion we are using this 

		fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)
	//by using <- in the fuction it will make the channel RECIVE ONLY otherwise channel is bidirectional
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mych <- 5
		//mych<-6
		wg.Done()
		defer close(mych) //use close after you have used your channel

	}(mych, wg)
	wg.Wait()
}
