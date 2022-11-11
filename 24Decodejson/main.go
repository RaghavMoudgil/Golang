package main

import (
	"encoding/json"
	"fmt"
)
type student struct {
	// Json format against these struct will work as their alias
	Name string 		//`json:"Firstname"`
	class string
	rollno int			 //`json:"-"`//You write - the data will not show up it can be used for passwords
	university string
}

func main() {
	fmt.Println("This code will help us to convert json Data into Golang")
	decodeJson()
}
func decodeJson()  { //this Json data is from the last program
	//as we have done in previous program the data from json is in byte format and we have 
	// to convert that data into string format here we have to do the opposite  
	jsonfromweb:= []byte(`
	{
		"name":"Raghav",
		"class":"Masters",
		"rollno":"454",
		"university":"KUK",
	}
		`)

	//here we are making a structure and going to fill it with the json data as we have done opposite in previous proram
	var studentData student

	//checking if the json is valid or not 
	checkvalid:=json.Valid(jsonfromweb)

	if checkvalid {
		fmt.Println("Json is valid")
		//repeating opposite steps marshel to unmarshel
		json.Unmarshal(jsonfromweb, &studentData)
		fmt.Printf("%#v\n",studentData)//For these type of statements # is used 
		
	}else{
		fmt.Println("Their is some error while converting the JSON data")
	}
}