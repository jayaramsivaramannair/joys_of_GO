package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

//http.ResponseWriter is used to send the response back to the client
//http.Request is used to process the incoming request from the client
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Below code encodes the response as json
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Helps with accessing parameters in the query string
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			//Three dot is used to unpack each movie from the second slice and push it into the first slice
			movies = append(movies[:index], movies[index+1:]...)
			break //if the specific movie is found then break out of the loop
		}
	}

	//This will return all the movies except the deleted one
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
	
}

func main() {
	//Helps create a new router which available on the mux package
	r := mux.NewRouter()


	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title : "Movie One", Director : &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title : "Movie Two", Director : &Director{FirstName: "Steve", LastName: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")

	//Listen and Serve function is used to start the server
	log.Fatal(http.ListenAndServe(":8000", r))


}

