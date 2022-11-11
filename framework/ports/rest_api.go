package framework

import "net/http"

type RESTPort interface {
	Run(string)
	RESTCreateMovie(http.ResponseWriter, *http.Request)
	RESTDeleteMovie(http.ResponseWriter, *http.Request)
	RESTGetAllMovies(http.ResponseWriter, *http.Request)
	RESTGetMovie(http.ResponseWriter, *http.Request)
	RESTUpdateMovie(http.ResponseWriter, *http.Request)
}
