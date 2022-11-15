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

	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	csvFile, err := os.Create("Data2.csv")
	if err != nil {
		panic(err)
	}
	//inserting data in CSV file
	writer := csv.NewWriter(csvFile)

	for _, line := range data {
		student := studData{
			Name:   line[0],
			Rollno: line[1],
			Class:  line[2],
		}
		fmt.Println(student.Name, " ", student.Rollno, " ", student.Class)
		_ = writer.Write(line)
	}

	fmt.Println("Your data is generate in new csv file ")
	//creating new CSV file
	writer.Flush()
	defer csvFile.Close()
}
