package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() *sql.DB {
	var err error
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/student?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}

type studData struct {
	name    string
	rollno  string
	class   string
	hindi   string
	english string
	maths   string
	total   string
}

func main() {
	fmt.Println("Code to create a dynamic structure..... in progress...")
	var db = initDB()
	row, err := db.Query("select * from student")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var temp []string
	var temp2 [][]string
	stud := studData{}
	for row.Next() {
		err = row.Scan(&stud.name, &stud.rollno, &stud.class, &stud.hindi, &stud.english, &stud.maths, &stud.total)
		if err != nil {
			log.Fatal(err)
		}
		//converting data into 2d slice by appending

		temp = []string{stud.name, stud.rollno, stud.class, stud.hindi, stud.english, stud.maths, stud.total}
		temp2 = append(temp2, temp)
	}
	defer fmt.Println(temp2)
}
