package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopher-camp/pkg/routes"
	"log"
	"net/http"
)

func main() {
	port := ":8008"
	r := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))

	r.Handle("/", fileServer)
	routes.RegisterProjectRoutes(r)
	http.Handle("/api/projects/", r)

	fmt.Printf("Starting server on port " + port + "\n")
	log.Fatal(http.ListenAndServe(port, r))
}
