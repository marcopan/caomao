package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)
/**
git test
**/
func main(){
	fmt.Println("test")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test",TestHandler/*1*/).Methods("GET")

	var address = fmt.Sprintf("%s:%d","localhost",8080)
	error := http.ListenAndServe(address,router)
	if(error != nil){
		log.Fatalln("error ",error)
	}
	fmt.Println("Server end")
}

func TestHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Write([]byte("test"))
}