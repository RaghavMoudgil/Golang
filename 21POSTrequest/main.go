package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("This is a code to handel POST request.")
	performPostJsonRequest()
}

// Generating a fake JSON data or payload
func performPostJsonRequest() {
	const myurl = "http://localhost:3000/post"
	requestbody := strings.NewReader(`
	{
		"Name":"Raghav",
		"Rollno":"654",
		"class":"Masters"
	}`)

	response, err := http.Post(myurl, "application/json", requestbody) //we can see in the thundercloud server header name is appllication and we want that
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
