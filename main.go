package main

import (
	"ProjectAssignment/router"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	router.SetupRoutes(r)

	// Start the server
	port := ":8080"
	fmt.Printf("Server is running on %s\n", port)
	http.Handle("/", r)
	http.ListenAndServe(port, r)
}
