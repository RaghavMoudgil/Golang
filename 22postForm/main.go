package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("This code will help to learn about post form.")
	PerformPostJsonFormRequest()
}
func PerformPostJsonFormRequest()  {
		const myurl= "http://localhost:3000/postform"
		//adding data to postform
		data:=url.Values{}
			data.Add("Name","Raghav")
			data.Add("class","masters")
			data.Add("rollno","64")

	response,err:=http.PostForm(myurl,data)		
	if err != nil {
		panic(err)
	}
	content,err:=ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(content))
}