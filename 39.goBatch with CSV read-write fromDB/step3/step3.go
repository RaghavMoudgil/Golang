package step3

import (
	"fmt"
	"student/connfunc"
)

func Clear_table() {
	var db = connfunc.InitDbStud()
	db.Exec("DELETE FROM student WHERE Total_marks = ''")
	fmt.Println("Data Cleared!!")
}
