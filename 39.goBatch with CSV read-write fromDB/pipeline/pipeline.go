package pipeline

import (
	"student/processor"
	"student/reader"
	// "student/step3"
	"student/task"
	"student/writer"

	"github.com/chararch/gobatch"
	_ "github.com/go-sql-driver/mysql"
)

var Counter int = 0

func JobRunner() string {
	//build steps
	step1 := gobatch.NewStep("mytask").Handler(task.Task).Build()
	step2 := gobatch.NewStep("my_step").Reader(&reader.MyReader{}).Processor(&processor.MyProcessor{}).Writer(&writer.MyWriter{}).ChunkSize(10).Build()
	// step3 := gobatch.NewStep("clear_table").Handler(step3.Clear_table).Build()
	//build job
	job := gobatch.NewJob("my_job").Step(step1, step2, step1, step2).Build()

	//register job to gobatch
	gobatch.Register(job)
	Counter++
	return job.Name()

}
