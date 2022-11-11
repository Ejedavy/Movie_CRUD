package framework

import "net/http"

type RESTPort interface {
	Run(string)
	CreateMovie(response http.ResponseWriter, request *http.Request)
	DeleteMovie(response http.ResponseWriter, request *http.Request)
	GetAllMovies(response http.ResponseWriter, request *http.Request)
	GetMovie(response http.ResponseWriter, request *http.Request)
	UpdateMovie(response http.ResponseWriter, request *http.Request)
}
