package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopher-camp/pkg/handlers"
	"log"
	"net/http"
)

func fahdHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "fahdjamy" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	_, err := fmt.Fprintf(w, "welcome")
	if err != nil {
		http.Error(w, "500", http.StatusBadRequest)
		return
	}
}

var projects []handlers.Project

func main() {
	port := ":8008"
	r := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))
	//
	r.Handle("/", fileServer)
	//http.HandleFunc("/fahd", fahdHandler)
	//

	r.HandleFunc("/projects", handlers.GetProjects).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.GetOneProject).Methods("GET")
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", handlers.UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id}", handlers.DeleteProjects).Methods("DELETE")

	fmt.Printf("Starting server on port " + port + "\n")
	log.Fatal(http.ListenAndServe(port, r))
}
