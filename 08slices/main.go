package main
import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is the code for slices")
	var fruitList = []string{} //We cannnot use below used method make here as another append will show error values will are not right 
	fmt.Printf("Type of fruitlist is : %T \n",fruitList)
	//Adding values in slice
	fruitList=append(fruitList, "mango","banana","peach","guavava","cherry","strawberry")
	fmt.Println(fruitList)
	//colon (:) is used to slice up the slice
	fruitList = append(fruitList[1:]) // this will remove the first member
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:5])// This will indicate first and last position & all in between is the output 
	fmt.Println(fruitList)


	
	scores:= make([]int, 4)
	scores[0]=111
	scores[01]=112
	scores[02]=115
	scores[03]=145
	//scores[4]=654   //this will show an error as memory space is 4
	fmt.Println(scores)
	// This will not show an error due to append method this will realocate the memory 
	scores= append(scores, 63541,351,51,3151,6541)
	fmt.Println(scores)
	fmt.Println("check weather ints are sorted or not:\n Ints are sorted", sort.IntsAreSorted(scores))
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println("check weather ints are sorted or not:\n Ints are sorted", sort.IntsAreSorted(scores))

	//Removing single or multiple value from the slice 

	car :=[]string{"verna","honda city","swift","nano","Beemer"}
	fmt.Println(car)
	//now if i have to remove nanp
	index := 1
	car =append(car [:index],car[index+1:]... )
	fmt.Println(car)
	// now if i have to remove multiple elements 
	car1 :=[]string{"verna","honda city","swift","nano","Beemer"}
	car1 =append(car1 [:index],car1[index+3:]... ) // this is concatinating two slices 
	fmt.Println(car1)

	//another way to remove multiple elements 
	car2 :=[]string{"verna","honda city","swift","nano","Beemer"}
	car2 =append(car2[1:4])// this will only show in between element.
	fmt.Println(car2)

}