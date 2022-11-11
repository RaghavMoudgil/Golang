package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Global variable for URL
const url ="https://google.com"
func main() {
	fmt.Println("This code will help us to learn how to deal with the web requests \n ")
	//package we are going to use is net/http
	responce,err:=http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response if of type :%T\n",responce)

	//alaways close the connection
	defer responce.Body.Close()

	//reading responce 
	databyte,err:= ioutil.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("This is the response body:",string(databyte))
}