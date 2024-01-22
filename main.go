package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihir010/ylytic/controllers"
)

func main() {
	r := mux.NewRouter()

	port := "8080"

	r.HandleFunc("/", controllers.ServeHome)

	// r.HandleFunc("/all", controllers.GetAll).Methods("GET")
	r.HandleFunc("/search", controllers.SearchParams).Methods("GET")

	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Server running on port: " + port)
}
