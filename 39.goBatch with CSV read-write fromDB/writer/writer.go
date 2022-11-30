package writer

import (
	"fmt"
	"log"
	"student/connfunc"
	"student/module"

	"github.com/chararch/gobatch"
	_ "github.com/go-sql-driver/mysql"
)

// writer
type MyWriter struct {
}

func (r *MyWriter) Write(item []interface{}, chunkCtx *gobatch.ChunkContext) gobatch.BatchError {
	var tmarks []string

	for _, v := range item {
		tmarks = v.([]string)
	}
	// fmt.Println("value", tmarks, reflect.TypeOf(tmarks))

	var db = connfunc.InitDbStud()
	var roll []string
	_ = roll
	rows, err := db.Query("SELECT  Rollno FROM student")
	if err != nil {
		fmt.Print(err)
	}
	// fmt.Print(rows)
	var stud module.StudData
	for rows.Next() {
		err = rows.Scan(&stud.Rollno)
		roll = append(roll, stud.Rollno)
		if err != nil {
			log.Fatal(err)
		}

	}
	// fmt.Print("stud roll", roll)

	for i, _ := range roll {
		db.Exec("UPDATE student SET Total_marks = ? WHERE Rollno = ? ", tmarks[i], roll[i])
	}

	defer fmt.Println("Processed data Written successfully")

	return nil
}
