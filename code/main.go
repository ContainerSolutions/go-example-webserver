package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Golang"))
}
