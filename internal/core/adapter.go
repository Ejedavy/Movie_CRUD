package core

import (
	"math/rand"
	"strconv"
)

type MovieAdapter struct {
}

func NewMovieAdapter() *MovieAdapter {
	return &MovieAdapter{}
}

func (adapter MovieAdapter) Create(movie Movie) Movie {
	movie.ID = strconv.Itoa(rand.Int())
	return movie
}
