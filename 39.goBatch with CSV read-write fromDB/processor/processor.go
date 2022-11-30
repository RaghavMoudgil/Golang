package processor

import (
	"log"
	"strconv"
	"student/connfunc"
	"student/module"

	"github.com/chararch/gobatch"
	_ "github.com/go-sql-driver/mysql"
)

// processor
type MyProcessor struct {
}

func (r *MyProcessor) Process(item interface{}, chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	var db = connfunc.InitDbStud()
	row, err := db.Query("select Hindi_marks,English_marks,Maths_marks from student")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var temp []string
	stud := module.StudData{}

	for row.Next() {
		err = row.Scan(&stud.Hindi_marks, &stud.English_marks, &stud.Maths_marks)
		if err != nil {
			log.Fatal(err)
		}
		int1, _ := strconv.Atoi(stud.English_marks)
		int2, _ := strconv.Atoi(stud.Hindi_marks)
		int3, _ := strconv.Atoi(stud.Maths_marks)

		stud.Total_marks = strconv.Itoa(int1 + int2 + int3)

		temp = append(temp, stud.Total_marks)
	}
	// fmt.Println("CHECKPOINT", temp)
	return temp, nil

}
