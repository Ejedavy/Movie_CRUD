package main

import (
	"github.com/ejedavy/movie_crud/database"
	framework "github.com/ejedavy/movie_crud/framework/ports"
	"github.com/ejedavy/movie_crud/internal/application"
	"github.com/ejedavy/movie_crud/internal/core"
	restServ "github.com/ejedavy/movie_crud/restAdapter"
)

func main() {
	var db framework.DatabasePort
	var restServer framework.RESTPort
	var api application.FrameworkInterface
	var app core.MovieInterface

	db = database.NewFileDB()
	app = core.NewMovieAdapter()
	api = application.NewAPIAdapter(db, app)
	restServer = restServ.NewRestServer(api)
	restServer.Run(":8000")
}
