package framework

import "github.com/ejedavy/movie_crud/internal/core"

type DatabasePort interface {
	Store(core.Movie) (core.Movie, error)
	DeleteMovie(string) ([]core.Movie, error)
	GetMovie(string) (core.Movie, error)
	GetAllMovies() ([]core.Movie, error)
}
