package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	//mod init command helps us to build the .exe file from our code 
	fmt.Println("This code will help us to learn about MOD init command")
	greater()
	r :=  mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")

	//running the server
	log.Fatal(http.ListenAndServe(":4000", r))
	
}
func greater()  {
	fmt.Println("Hello MUX user")
}
func serveHome(w http.ResponseWriter ,r *http.Request)  {
	//W stands for response writer (sending)
	//R stands for writer (receving)
	w.Write( []byte(`<h1>Welcome to Golang</h1> <CENTER><IMG SRC='C:\Users\raghavm\Desktop\Wallpaper\1045045.png' ALIGN='BOTTOM'> </CENTER>
	<HR>
	
	<a href=''>Link Name</a>
	
	is a link to another nifty site
	
	<H1>This is a Header</H1>
	
	<H2>This is a Medium Header</H2>
	
	Send me mail at <a href=''>
	
	support@yourcompany.com</a>.
	
	<P> This is a new paragraph!
	
	<P> <B>This is a new paragraph!</B>
	
	<BR> <B><I>This is a new sentence without a paragraph break, in bold italics.</I></B>
	
	<HR>`))

}