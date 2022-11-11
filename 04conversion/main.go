package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Welcome!! \n What is your number?")
	 
	reader:=bufio.NewReader(os.Stdin)
	input,_:=reader.ReadString('\n')
	fmt.Println("this is the number",input)

//	newname= input+ lastname          this error can be resolved by 

	newnumber, err := strconv.ParseFloat(strings.TrimSpace(input),64) //64 is the bit size

	// if you are going to use err then
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("addition of 1 is done :" ,newnumber+1)
	}

}