package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jeremydwayne/FEM_movies/models"
)

type MovieHandler struct{}

func (m *MovieHandler) writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (m *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			Title:       "The Matrix",
			ReleaseYear: 1999,
			Genres: []models.Genre{
				{ID: 1, Name: "Action"},
				{ID: 2, Name: "Sci-Fi"},
			},
			Keywords: []string{},
			Casting:  []models.Actor{{ID: 1, FirstName: "Keanu", LastName: "Reeves"}},
		},
		{
			ID:          2,
			Title:       "Back to the Future",
			ReleaseYear: 1984,
			Genres: []models.Genre{
				{ID: 1, Name: "Action"},
				{ID: 2, Name: "Sci-Fi"},
			},
			Keywords: []string{},
			Casting:  []models.Actor{{ID: 2, FirstName: "John", LastName: "Doe"}},
		},
	}

	m.writeJSON(w, movies)
}
