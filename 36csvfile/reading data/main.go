package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Rading data from the csv file.")

	//Opening a CSV file
	file, err := os.Open("Data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	//loop for reading the data
	for {
		rec, err := reader.Read()
		if err != nil {
			panic(err)
		}
		fmt.Println(rec)
	}

}
