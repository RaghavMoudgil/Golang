package main

import "fmt"

func main() {
	fmt.Print("This is the Mapping code\n")
	//You can use  make to create maps as well 
	languages:=make(map[string]string)
	languages["rb"]="ruby"
	languages["c#"]="c sharp"
	languages["py"]="python"
	languages["js"]="javascript"
	fmt.Println("These are all lannguages with abriviation:", languages)
	//Getting single value 
	fmt.Println(languages["js"])
	//Deleting a value from maps 
	delete(languages,"rb")
	fmt.Println(languages)

	//looping throughout the maps 

	
	for key, value := range languages{
		fmt.Println(key,value)
	}
}