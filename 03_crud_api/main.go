package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/manish39x/crud_api/model"
)

var movies []model.Movie

func main() {
	router := mux.NewRouter()

	movies = append(movies, model.Movie{ID: "1", Isbn: "346234", Title: "Ra one", Director: &model.Director{Firstname: "Rohit", Lastname: "shetty"}})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Lets get started with our crud api"))
	})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("SERVER IS LISTNING ON PORT: 8080")
	http.ListenAndServe(":8080", router)
}

// CONTROLLERS
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie model.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var movie model.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))

	for i, item := range movies {
		if item.ID == params["id"] {
			fmt.Println("enters")
			movies = slices.Delete(movies, i, i+1)
			movies = append(movies, movie)
		}
	}
	json.NewEncoder(w).Encode(movie)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = slices.Delete(movies, index, index+1)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
