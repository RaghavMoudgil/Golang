package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Creating a temporary sturcuture to import the data from the CSV file
type studData struct {
	Name   string
	Rollno string
	Class  string
}

func main() {

	fmt.Println("Rading data from the csv file.")

	//Opening a CSV file
	file, err := os.Open("Data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//Reading the data from the CSV file
	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	//Creating a new CSV file
	csvFile, err := os.Create("Data2.csv")
	if err != nil {
		panic(err)
	}
	//inserting data in new CSV file
	writer := csv.NewWriter(csvFile)

	for _, line := range data {
		student := studData{
			Name:   line[0],
			Rollno: line[1],
			Class:  line[2],
		}
		fmt.Println(student.Name, " ", student.Rollno, " ", student.Class)
		//using the same loop to write the data in new CSV file
		_ = writer.Write(line)
	}

	fmt.Println("Your data is generate in new csv file ")
	//creating new CSV file
	writer.Flush()
	defer csvFile.Close()
}
