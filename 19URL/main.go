package main

import (
	"fmt"
	"net/url"
)

// crearting a hypothetical url
const myurl string="https://www.youtube.com:8080/learn?coursename=reactjs&watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"
func main() {
	fmt.Println("This code will help you to learn about the URL\n ")
	//smae net/http pakage we are going to use 
	result,err:=url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Host)
	fmt.Println(result.User)
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)

	//stroing all the data of the prameter in query 
	qparam:=result.Query()
	fmt.Printf("the type of query param is:%T\n",qparam)
	//now qparam is the object which contains key value pair of all the URL data
	fmt.Println(qparam["coursename"],qparam["list"])//check from the url above

	//applying loop to get all the values in qparam
	for key, value := range qparam {
		fmt.Println("param id:",key,value)
		
	}

	//constructing a URL
	partsOfUrl :=&url.URL{
		Scheme: "https",            //remmember the way its written
		Host: "something.com",
		Path: "/somwone",
		RawPath: "somethind:someone",
	}
	anotherUrl:=partsOfUrl.String()
	fmt.Println(anotherUrl)
}