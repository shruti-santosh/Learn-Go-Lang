package main

import (
	"encoding/json"
	"fmt"
	"log"
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

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)	
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range(movies) {
		if item.ID == params["id"] {
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, item := range(movies) {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "4561", Title: "Title 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "7879", Title: "Title 2", Director: &Director{Firstname: "Dwayne", Lastname: "Johnson"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovieById).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8008\n")
	log.Fatal(http.ListenAndServe(":8008", r));
}