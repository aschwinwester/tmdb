package main

import (
	"tmdb/server"
	"tmdb/repository"
)

func main() {
	

	movieRepository := &repository.MRepository
	server := server.Server{MovieRepository: movieRepository}
	server.Startup()
}