package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jayaramsivaramannair/joys_of_GO/Project_CRUD-API-With-GORM/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	
}