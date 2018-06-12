package server

import (
	"fmt"
	"tmdb/model"
	"log"
	"net/http"
	"tmdb/repository"
	"encoding/json"

)
type Server struct {
	MovieRepository *repository.MovieRepository
}

func (s Server) Startup() {
	http.HandleFunc("/search", s.HandlerSearchActor)
	http.HandleFunc("/movies", s.HandlerGetMovies)
	http.ListenAndServe(":8080", nil)
}
func (s Server) HandlerGetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	movies := s.MovieRepository.GetMovies()
	b, err := json.Marshal(movies)

	if (err != nil) {
		log.Fatal(err)
	}
	w.Write(b)
}

func (s Server) HandlerSearchActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	v := r.URL.Query();
	firstname := v.Get("firstname")
	character := v.Get("character")
	if (firstname != "") {
		actors := s.MovieRepository.SearchActorsByFirstName(firstname)
		b, _ := json.Marshal(actors)
		w.Write(b)
	} else if (character != "") {
		actors := s.MovieRepository.SearchActorsCharacter(character)
		movies := s.MovieRepository.GetMovies()
		newmovies := make([]model.Movie, 0)
		for _ , actor := range actors {
			fmt.Printf("acteur %s met id %d\n", actor.Name, actor.Id)
			for _, movie := range movies {
				if (movie.IsActorInMovie(actor)) {
					newmovies = append(newmovies, movie)
				}
			}
		}
		b, _ := json.Marshal(newmovies)
	    w.Write(b)
		
	}
	
}