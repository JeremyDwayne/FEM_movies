package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jeremydwayne/FEM_movies/data"
	"github.com/jeremydwayne/FEM_movies/logger"
)

type MovieHandler struct {
	Storage data.MovieStorage
	Logger  *logger.Logger
}

func (m *MovieHandler) writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		m.Logger.Error("JSON Encoding error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (m *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := m.Storage.GetTopMovies()
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		m.Logger.Error("Unable to retrieve Top Movies", err)
	}

	m.writeJSON(w, movies)
}

func (m *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := m.Storage.GetRandomMovies()
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		m.Logger.Error("Unable to retrieve Random Movies", err)
	}

	m.writeJSON(w, movies)
}
