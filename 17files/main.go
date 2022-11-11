package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("code to unserstand files in Go lang ")
	content:= "This is the text that will go into the created file by the OS package"
	//file and err are used to show the file is created or it has some error 
	file,err:=os.Create("./GolangFile.txt")
	if err != nil {
		panic(err)
	}
	//IO package is used to write in the file
	length,err :=io.WriteString (file,content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is :",length)
	defer file.Close()
	readFile("./GolangFile.txt")
}
func readFile(filename string){
	//ioutil package is ued to read file data 
	databyte,err:=ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	//if we do not read file in string it will return us numbers or JSON format data 
	fmt.Println("The data in the file is",string(databyte))
}