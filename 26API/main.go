package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"math/rand"
	"github.com/gorilla/mux"
)

// creating model for work (in this case COURSE)
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int  `json:"courseprice"`
	Author      *Author `json:"author"` // this datatype is selfdefined as in author type is author

}

type Author struct {
	Fullname string `json:"fullname"` //use first letter as you are going to use it publicaly
	Website  string `json:"website"`
}

// creating a pysudodatabse
var courses []Course

// can call these as middleware or can also be called as helper file
func (c *Course) IsEmpty() bool { //c *Course is the method in the function IsEmpty
	//This case is used if we want to get course id from the user
	//  return c.CourseId == "" && c.CourseName == ""
	//if we want generate a random course id we will do this instead
	return c.CourseName == ""
}

// now we will work on controller
// servehome route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the API </h1>"))
}

/* #region controlling Part */
// sending all the values from the databse
func getAllData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all data of courses")
	//setting headers
	w.Header().Set("Content-type", "application/json")
	//throughing up all the things in as a json
	json.NewEncoder(w).Encode(courses)

}

// sending one value from the databse
func getonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content-Type", "applicartion/json") //setting header
	//grabbing the unique ID of the course
	params := mux.Vars(r) //rather than using URL as we have done earlier we are using mux here
	//which will allow us to get the variable usig Vars method from the request
	fmt.Println(params)

	//now we will loop through all the courses and we will going to find the matching ID from the request
	for _, course := range courses {
		if course.CourseId == params["id"] { //we will get to know about this "id" when we come onto this routing part
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Id if found here that you provided")
	return

}

func creatOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one Course")
	w.Header().Set("Course-Type", "application/json")

	//In case of empty request body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some Data")
	}
	//In case they send empty data
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data found")
		return
	}
	
	//Now we will automatically going to generate a courseID
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode("Course Successfully Added!! ")
	json.NewEncoder(w).Encode(course)
}

// updating Courses
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course ")
	w.Header().Set("Content-Type", "applicatoin/json")

	//grabbing Id from request
	params := mux.Vars(r)

	//now we are going to get the id from the body
	//far that we are going to loop through the body to find the ID after that
	//we will remove that value from the index
	//after that we will be going to add the value of the id which we are reciving from the request (params)
	// here we are going to use param 2 times first for finding the ID and second for adding

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			//as json data is coming to us we are going to use a DECODER here
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("value Updarted!!!")
			return
		}
	}
}

// deleting the course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a single course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//same we are going to do is loop through the index and then remove that index
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Value is Deleted!!")
			break
		}
	}
}

/*#endregion Ending of Controlling Part*/

func main() {
	fmt.Println("This cade is to learn about how to build API (CRUD) ")
	//in this we will only going to build Api data integration part will come in next code

	//In terminal create mod file using commmand /*go mod init <name_of _mod_file>*/

	//Use command /*go get -u github/gorilla/mux*/ to get mux file

	/*#region Routing Part*/

	r := mux.NewRouter()

	//seeding (simply means sending data into the courses slice)
	courses = append(courses, Course{
		CourseId: "2", CourseName: "javascript", CoursePrice: 654, Author: &Author{
			Fullname: "Raghav",
			Website:  "raghav.js.dev",
		},
	})
	courses = append(courses, Course{
		CourseId: "6", CourseName: "golang", CoursePrice: 20, Author: &Author{
			Fullname: "Raghav",
			Website:  "raghav.go.dev",
		},
	})



	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllData).Methods("GET")
	r.HandleFunc("/course/{id}", getonecourse).Methods("GET")
	r.HandleFunc("/course", creatOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


	//listning to the port
	log.Fatal(http.ListenAndServe(":3000", r))
	/*#endregion ending of Routing Part*/
}
