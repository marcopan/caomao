package main

import (
	"net/http"
)

func TestHandler1(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Write([]byte("test"))
}