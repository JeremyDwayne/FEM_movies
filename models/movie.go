package models

type Movie struct {
	ID          int      `json:"id"`
	TMDB_ID     int      `json:"tmdb_id"`
	Title       string   `json:"title"`
	Tagline     string   `json:"tagline"`
	ReleaseYear int      `json:"release_year"`
	Genres      []Genre  `json:"genre"`
	Overview    *string  `json:"overview"`
	Score       *float32 `json:"score"`
	Popularity  *float32 `json:"popularity"`
	Keywords    []string `json:"keywords"`
	Language    *string  `json:"language"`
	PosterURL   *string  `json:"poster_url"`
	TrailerURL  *string  `json:"trailer_url"`
	Casting     []Actor  `json:"casting"`
}
