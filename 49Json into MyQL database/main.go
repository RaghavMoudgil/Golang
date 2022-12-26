package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type studData struct {
	Name   string
	Rollno string
	Class  string
}

func initDB() *sql.DB {
	var err error
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/student?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	fmt.Println("In this code we will take data from the JSON file and going to write it in SQL DB.")
	//setting connection to the database
	var db = initDB()
	fmt.Println("Reading data from the JSON file.")
	//Opening a CSV file
	file, err := ioutil.ReadFile("data.json")

	if err != nil {
		log.Fatal(err)
	}
	//using Unmarshel to parse the JSON data
	data := []studData{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	//writing data into the Database
	for i, _ := range data {

		_, err = db.Exec("Insert into json_test (Name,Rollno,Class) values (?,?,?)", data[i].Name, data[i].Rollno, data[i].Class)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer fmt.Println("Data uploaded to SQL")

}
