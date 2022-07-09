package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	getMovies(w, r)
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
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, newMovie)
	getMovies(w, r)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = params["id"]
	for i, item := range movies {
		if item.ID == newMovie.ID {
			movies[i] = newMovie
			break
		}
	}
	getMovies(w, r)
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "23234324", Title: "Something", Director: &Director{Firstname: "Bugs", Lastname: "Bunny"}})
	movies = append(movies, Movie{ID: "2", Isbn: "25634324", Title: "Something else", Director: &Director{Firstname: "Daffy", Lastname: "Duck"}})
	// fmt.Print(movies)
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
