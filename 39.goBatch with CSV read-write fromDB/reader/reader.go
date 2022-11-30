package reader

import (
	"encoding/csv"
	"fmt"
	"os"
	"student/connfunc"
	"student/module"

	"github.com/chararch/gobatch"
	_ "github.com/go-sql-driver/mysql"
)

// reader
type MyReader struct {
}

func (r *MyReader) Read(chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	curr, _ := chunkCtx.StepExecution.StepContext.GetInt("read.num", 0)
	if curr < 100 {
		chunkCtx.StepExecution.StepContext.Put("read.num", curr+1)
		return fmt.Sprintf("value-%v", curr), nil
	}

	fmt.Println("Creating and Reading data from the File")
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
	newdata := []module.StudData{}
	_ = newdata
	for _, value := range data {
		newdata = append(newdata, module.StudData{Name: value[0], Rollno: value[1], Class: value[2], Hindi_marks: value[3], English_marks: value[4], Maths_marks: value[5], Total_marks: value[6]})
	}
	var db = connfunc.InitDbStud()
	for i, _ := range newdata {
		db.Exec("insert into student(Name,Rollno,Class,Hindi_marks,English_marks,Maths_marks,Total_marks) values (?,?,?,?,?,?,?)", newdata[i].Name, newdata[i].Rollno, newdata[i].Class, newdata[i].Hindi_marks, newdata[i].English_marks, newdata[i].Maths_marks, newdata[i].Total_marks)
	}
	fmt.Println("Data written to MYSQL database.")
	return nil, nil
}
