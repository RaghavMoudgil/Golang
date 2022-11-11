package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	Welcome:= "Welcome User"
	fmt.Println(Welcome)
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:" )
	input,_:=reader.ReadString('\n')
	fmt.Print("this is your name:",input )
	fmt.Printf("and type is : %T",input)
}