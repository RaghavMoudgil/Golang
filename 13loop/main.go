package main

import "fmt"

func main() {
	fmt.Println("Loops in Go lang")
	days:=[]string{"sunday","monday","tuesday","wednesday","thursday","friday","saturday"}
	fmt.Println(days)

	for i:=0;i<len(days);i++{
		fmt.Println(days[i])
	}
	for j:= range days{
		fmt.Println(days[j])
	}
	//Most userful
	for index, day :=range days{
		fmt.Printf("index is %v and value is %v\n",index,day)
	}
	//Using as a while loop 
	somevalue:=1
	for somevalue<20{
	fmt.Println(somevalue)
	somevalue++
	if somevalue==15{
		goto ft
	}
	}
	ft:
	fmt.Println("Comes from Goto statement")
}