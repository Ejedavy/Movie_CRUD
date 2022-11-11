package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (serv RESTServer) CreateMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	movie, _ := serv.api.CreateMovie(request.Body)
	json.NewEncoder(response).Encode(movie)
}
func (serv RESTServer) DeleteMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	movies, _ := serv.api.DeleteMovie(params["id"])
	json.NewEncoder(response).Encode(movies)
}
func (serv RESTServer) GetAllMovies(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	movies, _ := serv.api.GetAllMovies()
	json.NewEncoder(response).Encode(movies)
}
func (serv RESTServer) GetMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	movie, _ := serv.api.GetMovie(params["id"])
	json.NewEncoder(response).Encode(movie)
}

func (serv RESTServer) UpdateMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	movie, _ := serv.api.UpdateMovie(params["id"], request.Body)
	json.NewEncoder(response).Encode(movie)
}
