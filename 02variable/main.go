package main
import "fmt"
func main(){
	var username string ="Raghav"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n ", username  )
	
	var isLoggedin bool = true;
	fmt.Printf("this is boolean value: %T \n", isLoggedin  )

	var inteager uint =255 //it only takes upto 255 bigger than that is not valid 
	fmt.Println(inteager)
	fmt.Printf("this is the small inteager value: %T \n",inteager)

	var floating float32=255.321654987
	fmt.Print(floating)
	fmt.Printf("this is the float which allows only 5 decimal : %T ,\n", floating)

	var bigfloating float64=255.321654987
	fmt.Print(bigfloating)
	fmt.Printf("Float 64 will allow more than 2 decimal values: %T \n",bigfloating)
	
	var blue string
	fmt.Printf("finding what is in the smpty string :%T \n", blue )
	fmt.Println(blue) 
	// you can declair a variable wihtout specifing what type of variable it is 
	// you can also declair variable with the short-hand which is this 
	noVarVariable := 321654978  // this is only alowed inside the method not outside the method 
	fmt.Println(noVarVariable)
}