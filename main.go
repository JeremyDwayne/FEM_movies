package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/jeremydwayne/FEM_movies/data"
	"github.com/jeremydwayne/FEM_movies/handlers"
	"github.com/jeremydwayne/FEM_movies/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	// Logging
	logInstance := initializeLogger()

	// Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Fatal("DATABASE_URL NOT SET")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}

	defer db.Close()

	// Initialize Repositories
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize repository")
	}

	// Handlers
	movieHandler := handlers.MovieHandler{
		Storage: movieRepo,
		Logger:  logInstance,
	}

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	const addr = ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed", err)
	}
}
