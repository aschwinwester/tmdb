package main

import (
	"tmdb/server"
	"tmdb/repository"
)

// Create a repository from a file
// and creates a server. Then starts the server
func main() {
	movieRepository := &repository.MRepository
	server := server.Server{MovieRepository: movieRepository}
	server.Startup()
}