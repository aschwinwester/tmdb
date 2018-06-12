package repository
import (
	"log"
	"tmdb/model"
	"tmdb/parser"
	"os"
)
type MovieRepository struct {
	movies []model.Movie
}

func(m MovieRepository) GetMovies() []model.Movie {
	 return m.movies
}

func(m MovieRepository) SearchActorsByFirstName(name string) []model.Actor {
	actors := make([]model.Actor,0)
	for _, movie := range m.GetMovies() {
		as := movie.SearchActorsByFirstName(name)
		for _, actor := range as {
			actors = append(actors, actor)
		}
	}
	return actors;
}

func(m MovieRepository) SearchActorsCharacter(character string) []model.Actor {
	actors := make([]model.Actor,0)
	for _, movie := range m.GetMovies() {
		as := movie.SearchCharacter(character)
		for _, actor := range as {
			var found = false;
			for _, currentActor := range actors {
				if (currentActor.Name == actor.Name) {
					found = true
				}
			}
			if (!found) {
				actors = append(actors, actor)
			}
		}
	}
	return actors;
}

var MRepository MovieRepository = MovieRepository{}

func init() {
	csvFile, err := os.Open("tmdb_5000_credits.csv")
	if (err != nil) {
		log.Fatal(err)
	}
	defer csvFile.Close()
	
	MRepository = MovieRepository{movies: parser.ParseCsvFile(csvFile)}
}

