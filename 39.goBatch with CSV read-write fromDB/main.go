package main

import (
	"context"
	"student/connfunc"
	"student/step3"

	"student/pipeline"
	"time"

	"github.com/chararch/gobatch"
	"github.com/chararch/gobatch/util"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db = connfunc.InitDbJob()
	gobatch.SetDB(db)

	//run
	params, _ := util.JsonString(map[string]interface{}{
		"rand": time.Now().Nanosecond(),
	})
	step := gobatch.NewStep("clear_table").Handler(step3.Clear_table).Build()
	job := gobatch.NewJob("clearTable_job").Step(step).Build()
	gobatch.Register(job)
	gobatch.Start(context.Background(), pipeline.JobRunner(), params)
	// fmt.Println("Counter is ", pipeline.Counter)
	// gobatch.Start(context.TODO(), job.Name(), params)
}
