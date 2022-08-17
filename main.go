package main

import (
	"fmt"
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

func main() {
	port := ":8008"
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/fahd", fahdHandler)

	fmt.Printf("Starting server on port " + port + "\n")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
