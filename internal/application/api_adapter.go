package application

import (
	"encoding/json"
	"io"

	framework "github.com/ejedavy/movie_crud/framework/ports"
	"github.com/ejedavy/movie_crud/internal/core"
)

type APIAdapter struct {
	db  framework.DatabasePort
	app core.MovieInterface
}

func NewAPIAdapter(db framework.DatabasePort, app core.MovieInterface) *APIAdapter {
	return &APIAdapter{db: db, app: app}
}

func (adapter APIAdapter) CreateMovie(request io.Reader) (core.Movie, error) {
	var movie core.Movie
	json.NewDecoder(request).Decode(&movie)
	movie = adapter.app.Create(movie)
	movie, err := adapter.db.Store(movie)
	if err != nil {
		return core.Movie{}, err
	}
	return movie, nil
}

func (adapter APIAdapter) DeleteMovie(id string) ([]core.Movie, error) {
	var movies []core.Movie
	movies, err := adapter.db.DeleteMovie(id)
	if err != nil {
		return make([]core.Movie, 0), err
	}
	return movies, nil
}

func (adapter APIAdapter) UpdateMovie(id string, request io.Reader) (core.Movie, error) {
	var movie core.Movie
	json.NewDecoder(request).Decode(&movie)
	_, err := adapter.db.DeleteMovie(id)
	if err != nil {
		return core.Movie{}, err
	}
	movie = adapter.app.Create(movie)
	adapter.db.Store(movie)
	return movie, nil
}

func (adapter APIAdapter) GetAllMovies() ([]core.Movie, error) {
	return adapter.db.GetAllMovies()
}

func (adapter APIAdapter) GetMovie(id string) (core.Movie, error) {
	return adapter.db.GetMovie(id)
}
