package database

import (
	"errors"

	"github.com/ejedavy/movie_crud/internal/core"
)

type FileDB struct {
	movies []core.Movie
}

func NewFileDB() *FileDB {
	var movies = make([]core.Movie, 0)
	return &FileDB{movies: movies}
}

func (db *FileDB) Store(movie core.Movie) (core.Movie, error) {
	db.movies = append(db.movies, movie)
	return movie, nil
}

func (db *FileDB) DeleteMovie(id string) ([]core.Movie, error) {
	for index, mov := range db.movies {
		if mov.ID == id {
			db.movies = append(db.movies[:index], db.movies[index+1:]...)
			return db.movies, nil
		}
	}
	return make([]core.Movie, 0), errors.New("invalid id")
}

func (db *FileDB) GetMovie(id string) (core.Movie, error) {
	for _, mov := range db.movies {
		if mov.ID == id {
			return mov, nil
		}
	}
	return core.Movie{}, errors.New("invalid id")
}

func (db *FileDB) GetAllMovies() ([]core.Movie, error) {
	return db.movies, nil
}
