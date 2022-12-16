package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("This code is going to talk about the CONTEXt")
	//Context is created for each net/http requesty
	ctx := req.Context()
	fmt.Println("Server is now going to Start!!")
	defer fmt.Println("Server is now Closed!!")
	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("Server: ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
