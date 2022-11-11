package restServ

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
	log.Println("Registering Routes...")
	router.HandleFunc("/movie", serv.RESTCreateMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", serv.RESTDeleteMovie).Methods("DELETE")
	router.HandleFunc("/movie/{id}", serv.RESTUpdateMovie).Methods("PUT")
	router.HandleFunc("/movies", serv.RESTGetAllMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", serv.RESTGetMovie).Methods("GET")
	log.Println("Registered")
	log.Println("Starting http server at port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
