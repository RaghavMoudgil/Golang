package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	// Json format against these struct will work as their alias
	Name string `json:"Firstname"`
	class string
	rollno int `json:"-"`//You write - the data will not show up it can be used for passwords
}

func main() {
	fmt.Println("In this code we are going to create a Json Data in Golang")
	encodeJson()
}

// encoding of the JSON
func encodeJson() {
	studentData := []student{
		{"Raghav", "Masters", 321}, {"Suresh", "Phd", 3212}, {"Mahesh", "Bachelors", 4521},
	}
	//Now we will be packaging this data into Json data

	finalJson, err := json.MarshalIndent(studentData,"", "\t") //marshel is used to egt the json data and marshelIndent is used to make the json data more preety
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}
 