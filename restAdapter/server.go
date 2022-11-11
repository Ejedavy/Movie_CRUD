package server

import (
	"log"
	"net/http"

	"github.com/ejedavy/movie_crud/internal/application"
	"github.com/gorilla/mux"
)

type RESTServer struct {
	api application.FrameworkInterface
}

func NewRestServer(api application.FrameworkInterface) *RESTServer {
	return &RESTServer{api: api}
}

func (serv RESTServer) Run(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/movie", serv.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", serv.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/movies", serv.GetAllMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", serv.GetMovie).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
