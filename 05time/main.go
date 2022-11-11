package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This code is about time ")

	presenttime := time.Now()
	fmt.Println(presenttime)
	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday"))//Remmember this, this is the stander you have to remmeber it 
	createdDate:=time.Date(2020, time.April, 10, 5,5,0,0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday "))
	}	