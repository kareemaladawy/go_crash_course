package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateUserResponse struct {
	Greeting string `json:"greeting"`
}

type HelloResponse struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

type AboutResponse struct {
	About string `json:"about"`
	Count int    `json:"count"`
}

type Middleware func(http.Handler) http.Handler

func authenticateMW(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer secret" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HelloResponse{"Hello World", 1}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := AboutResponse{"About text", 2}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := CreateUserResponse{
		Greeting: fmt.Sprintf("Hello %s, you are %d years old", user.Name, user.Age),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/users", authenticateMW(http.HandlerFunc(createUserHandler)))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/about", about)

	fmt.Println("Server Listening..")
	http.ListenAndServe(":4001", mux)
}
