package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Welcome to modules in go lang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod user")
}

func serveHome(w http.ResponseWriter, r *http.Request) { //w ==> response writer r ===> request pointer
	w.Write([]byte("<h1>Welcome to go lang series on YT</h1>"))
}
