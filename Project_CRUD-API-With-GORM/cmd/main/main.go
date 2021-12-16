package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"github.com/jayaramsivaramannair/joys_of_GO/Project_CRUD-API-With-GORM/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}