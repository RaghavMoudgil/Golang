package main

import "fmt"

func main() {
	fmt.Println("This is the code for array ")
	var fruitlist [4]string
	fruitlist[0]="apple"
	fruitlist[1]="Grapes"
	fruitlist[3]="peach"
	fmt.Println("This is the list of fruit :",fruitlist)
	fmt.Println("length of the list is :", len(fruitlist))


	var vegitable = [4]string{"potato","tomato","bringel"}
	fmt.Println("This is the vegitable list ",vegitable)
	fmt.Println("This is the length of the vegitable list ",len(vegitable))

	// vegitable[0]="potato"
	// vegitable[1]="tomato"
	// vegitable[2]="egg plant"
	// vegitable[3]="mushroom"
}