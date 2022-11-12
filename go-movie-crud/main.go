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
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director `json:"director"`
}

type Director struct {
	firstName string `json:"firstname"`
	lastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, value := range movies {
		if params["id"] == value.ID {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
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

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Params := mux.Vars(r)
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	for index, item := range movies {
		if Params["id"] == item.ID {
			movies := append(movies[:index], movies[index+1:]...)
			movie.ID = Params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}

	}
}
func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", ISBN: "563", Title: "Superman", Director: Director{firstName: "Henry", lastName: "Cavil"}})
	movies = append(movies, Movie{ID: "2", ISBN: "334", Title: "Mission Imposible", Director: Director{firstName: "Daniel", lastName: "Craig"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
