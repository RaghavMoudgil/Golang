package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Creating a temporary sturcuture to import the data from the CSV file
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
	var db = initDB()

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
	newdata := []studData{}
	_ = newdata
	for _, value := range data {
		newdata = append(newdata, studData{Name: value[0], Rollno: value[1], Class: value[2]})
	}
	for i, _ := range newdata {
		db.Exec("insert into student(Name,Rollno,Class) values (?,?,?)", newdata[i].Name, newdata[i].Rollno, newdata[i].Class)
	}
	fmt.Println("Data written to MYSQL database.")

	//extracting data from the MYsql DB
	row, err := db.Query("select * from student")
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
		//converting data into 2d slice by appending

		temp = []string{stud.Name, stud.Rollno, stud.Class}
		temp2 = append(temp2, temp)
	}

	//now all the data is in temp2 variable and now we will going to print it in a CSV file
	//creating csv file
	csvfile, err := os.Create("DATA2.csv")
	if err != nil {
		log.Fatal(err)
	}
	//writing on the csv file
	writer := csv.NewWriter(csvfile)
	err = writer.WriteAll(temp2)
	if err != nil {
		fmt.Println(err)
	}

}
