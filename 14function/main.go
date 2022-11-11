package main

import "fmt"

func main() {
	fmt.Println("Using functions in Golang") //function inside the function is not allowed
	greet()
	result:= adder(23,234)
	fmt.Println("the sum is:",result)
	proResult,sumnumber :=proAdder(0,1,2,3,4,5,6,7,8,9)
	fmt.Printf("Pro sum is %v and sum number is %v\n",proResult,sumnumber)
}
func greet(){
	fmt.Println("Greetings golang user")
}
func adder(a int ,b int) int {   //the 3rd int said that what type of value we are going to retur or which type of value expected
	return a+b
}

//creating a function if you dont know the amount of argument you are going to pass

func proAdder(values...int )(int, int) {
total := 0
count:=0
 for index,value:= range values{
	total += value
	count+=index
 }
  return total,count
}
