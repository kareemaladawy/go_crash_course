package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About page")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Listening..")
	http.ListenAndServe(":4001", nil)
}
