package jobs

import (
	"batch/datareader"
	"batch/filehandler"
	"batch/pipeline"
	"batch/task"

	"github.com/chararch/gobatch"
)

func Job() string {

	//building steps
	step1 := gobatch.NewStep("task").Handler(task.Task).Build()
	step2 := gobatch.NewStep("my_step").Reader(&pipeline.Reader{}).Processor(&pipeline.Processor{}).Writer(&pipeline.Writer{}).ChunkSize(10).Build()

	//building JOB
	Job := gobatch.NewJob("my_job").Step(step1, step2).Build()

	// registering job to gobatch
	gobatch.Register(Job)
	return Job.Name()
}

func RunStudJob() string {
	step1 := gobatch.NewStep("import_studData").ReadFile(filehandler.ReadFile).Writer(&filehandler.SqlStudent{}).ChunkSize(10).Build()
	step2 := gobatch.NewStep("export_studData").Reader(&datareader.StudRreader{}).WriteFile(filehandler.CreateFile).ChunkSize(10).Build()
	step3 := gobatch.NewStep("Uploading_file_to_ftp").CopyFile(filehandler.CopyFileToFtp, filehandler.CopyChecksumFileToFtp).Build()

	job := gobatch.NewJob("file_creation_job").Step(step1, step2, step3).Build()

	gobatch.Register(job)
	return job.Name()
}
