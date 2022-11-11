package application

import (
	"io"

	"github.com/ejedavy/movie_crud/internal/core"
)

type FrameworkInterface interface {
	CreateMovie(io.Reader) (core.Movie, error)
	DeleteMovie(id string) ([]core.Movie, error)
	UpdateMovie(string, io.Reader) (core.Movie, error)
	GetAllMovies() ([]core.Movie, error)
	GetMovie(id string) (core.Movie, error)
}
