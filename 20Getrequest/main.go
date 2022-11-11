package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Code to get GET request from the server using GOlang")
	Getrequest()
}
func Getrequest(){
	const myurl = "http://localhost:3000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code:", response.StatusCode)
	fmt.Println("Content length is :", response.ContentLength)

	content,err:=ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content)) 
	//using strings pakage rather than string to get data from byte to json format
	//we will use this as it is more powerfull and it holds all the request data in a variable in our case which is bytecount 
	var responsestring strings.Builder
	bytecount,_:=responsestring.Write(content)
	fmt.Println("This is the bytecount of the response:",bytecount)
	//to get respomne we will do 
	fmt.Println(responsestring.String())
}