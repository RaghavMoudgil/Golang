package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("We are going to generate CSV file in this code.")
	//creating a 2D slice
	data := [][]string{
		{"Name", "Rollno", "Class"},
		{"Raghav", "45", "Masters"},
		{"Rahul", "654", "Phd"},
		{"suresh", "5", "10th"},
		{"Riya", "46", "Phd"},
	}

	//creating CSV file
	csvFile, err := os.Create("Data.csv")
	if err != nil {
		panic(err)
	}

	//inserting data in CSV file
	writer := csv.NewWriter(csvFile)

	//loop for writing file
	for _, row := range data {
		_ = writer.Write(row)

	}
	writer.Flush()
	csvFile.Close()

}
