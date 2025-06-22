package main

import (
	"log"
	"net/http"

	"github.com/jeremydwayne/FEM_movies/handlers"
	"github.com/jeremydwayne/FEM_movies/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {
	logInstance := initializeLogger()
	movieHandler := handlers.MovieHandler{}

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	const addr = ":8080"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		logInstance.Error("Server failed: %v", err)
	}
}
