package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	// Unbuffered channel (capacity 0)
	unbufChan := make(chan string)
	// Buffered channel (capacity 2)
	bufChan := make(chan string, 2)

	// Unbuffered channel demonstration
	go func() {
		fmt.Println("Sending to unbuffered channel (will block until received)...")
		unbufChan <- "Message 1 from unbuffered channel"
		fmt.Println("Sent Message 1 to unbuffered channel")
		unbufChan <- "Message 2 from unbuffered channel"
		fmt.Println("Sent Message 2 to unbuffered channel")
	}()

	time.Sleep(500 * time.Millisecond) // Ensure goroutine starts

	fmt.Println("Receiving from unbuffered channel...")
	fmt.Println(<-unbufChan)
	fmt.Println(<-unbufChan)

	// Buffered channel demonstration
	fmt.Println("Sending to buffered channel (won't block until buffer is full)...")
	bufChan <- "Message 1 from buffered channel"
	fmt.Println("Sent Message 1 to buffered channel")
	bufChan <- "Message 2 from buffered channel"
	fmt.Println("Sent Message 2 to buffered channel")

	go func() {
		fmt.Println("Receiving from buffered channel in goroutine...")
		bufChan <- "Message 3 from buffered channel"
		fmt.Println(<-bufChan)
		bufChan <- "Message 3 from buffered channel"
		fmt.Println(<-bufChan)
		fmt.Println(<-bufChan)

	}()
	//below is  the code tfor selecct statement
	//as soon as data recived from channels we get to know in the select
	select {
	case msg1 := <-unbufChan:
		fmt.Println("Received from unbuffered channel:", msg1)
	case msg2 := <-bufChan:
		fmt.Println("Received from buffered channel:", msg2)
	}

	time.Sleep(1 * time.Second) // Wait for goroutines to finish
	fmt.Println("Main finished")
}
