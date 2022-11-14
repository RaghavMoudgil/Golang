package main

import (
	"fmt"
	"time"
)

func main() {
	go grater("Hello")  //here program is not waiting for thr controller to come back so it is not able to execute the command
	grater("World")
}
func grater(s string) {
	for i := 0; i <= 5; i++ {
		//we are here letting program to wait for the execution of Go routine {hello}
		time.Sleep(2* time.Second) //now here program will wait for 2 sec befire every execution
		//but above is not the ideal way and not used in real world 
		// we will going to use mutex for that 
		fmt.Println(s)
	}
}