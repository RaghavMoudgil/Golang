package main

import (
	"fmt"
	"log"
	"mongoDB/router"
	"net/http"
)

func main() {
	fmt.Println("In this code we are going to learn about Mongodb setup for API")
	fmt.Println("Serever is Starting....")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8082", r))
	fmt.Println("Listning at port 4000 ...")

}
