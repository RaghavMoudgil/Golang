package main

//we are using go routines which are managed by go runtime, to impliment various threads at same time 
//but there is no one to manage the memory for the threads, Here mutex comes in between to manage the memry
//what mutex will do it will apply lock on the memory once it is used so that other runtime routine will not try to access it


import (
	"fmt"
	"net/http"
	"sync"
)
//geenerally we use pointer in waitgroup 
var wg sync.WaitGroup //we are usnig this to add wait in go routines 
var mut sync.Mutex //declairing mutex
var signals =[]string{"test"} //thios var is for testing mutex

func main() {
	fmt.Println("In this code we are going to learn about the MUTEX")
	wesiteList := []string{
		"https://github.com",
		"https://google.com",
		"http://facebook.com",
	}
	for _, web := range wesiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Print(signals)
}
// this all will take time we can do it a lot faster by launching different threads at same time

//so we are going to use Go keyword before the method so that it can launch multiple threads 
// but it will not going to wait, as their is no method for waiting in the main function 
// and will not show any output.
func getStatusCode(endpoint string) {
	defer wg.Done()
	
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code is %s \n", res.StatusCode, endpoint)

}
