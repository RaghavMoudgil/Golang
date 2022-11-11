package main

import (

	"fmt"

)

func main() {
	fmt.Println("If-Else in Golang")
	fmt.Println("Inpuir Marks")
	
	var result string
	if marks:=19; marks<=33{
		result="Student is Fail"
	}else if marks>=33 && marks<=75{
		result="Average Student"
	}else{
		result="Gud Student"
	}
		

	fmt.Println(result)
}