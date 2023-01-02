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
	fmt.Println("In this code we will retrive data from the SQL database and print it in the JSON format.")

	fmt.Println("Reading data from the Database.")
	var db = initDB()

	row, err := db.Query("select * from json_test")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var temp []string
	var temp2 [][]string
	stud := studData{}
	for row.Next() {
		err = row.Scan(&stud.Name, &stud.Rollno, &stud.Class)
		if err != nil {
			log.Fatal(err)
		}

		temp = []string{stud.Name, stud.Rollno, stud.Class}
		temp2 = append(temp2, temp)

	}
	data, err := json.Marshal(temp2)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("json_data.json", data, 100)
	if err != nil {
		log.Fatal(err)
	}

}
