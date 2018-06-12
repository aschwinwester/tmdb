package parser

import (
	"tmdb/model"
	"encoding/json"
	"encoding/csv"
	"os"
	"fmt"
)

func parseActors(line string) []model.Actor {

	var actors []model.Actor
	json.Unmarshal([]byte(line), &actors)
	return actors
}
func parseMovie(fields []string) model.Movie {
	movie := model.Movie{Id: fields[0], Title: fields[1], Actors: parseActors(fields[2])}
	return movie
}

func ParseCsvFile(file *os.File) []model.Movie {
	reader := csv.NewReader(file)
	
	records, err := reader.ReadAll()
	if (err !=nil) {
		fmt.Print(err)
	}

	fmt.Printf("nr or records found is %d\n", len(records))
	movies := make([]model.Movie, 0)
	for index, record := range records {
		if (index > 0) {
			movie := parseMovie(record)
			movies = append(movies, movie)
		}
	}
	return movies
}