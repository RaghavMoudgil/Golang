package pipeline

import (
	"fmt"

	"github.com/chararch/gobatch"
)

// here we are goin to write function for reader
type Reader struct {
	
}

func (r *Reader) Read(chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	curr, _ := chunkCtx.StepExecution.StepContext.GetInt("read.num", 10)
	if curr < 10 { //this is the read and write count
		chunkCtx.StepExecution.StepContext.Put("read.num", curr+1)
		return fmt.Sprintf("value-%v", curr), nil
	}

	//data := []string{"mango", "apple", "banana"}
	return nil, nil
}

// here we are going to write function for processor
type Processor struct {
}

func (r *Processor) Process(item interface{}, chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	// var fruits []string
	// for k, v := range item.([]string) {
	// 	if k == 2 {
	// 		fmt.Println("data from previous step", v)
	// 		fruits = append(fruits, v)
	// 	}
	// 	fmt.Println("data after processing", fruits)

	// }

	return fmt.Sprintf("processec-%v", item), nil
}

// here we are going to write function for writer
type Writer struct {
}

func (r *Writer) Write(items []interface{}, chunkCtx *gobatch.ChunkContext) gobatch.BatchError {

	fmt.Printf("Write: %v\n", items)
	return nil
}
