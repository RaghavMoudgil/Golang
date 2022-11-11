package router

import (
	"mongoDB/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMoviesController).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("api/movies/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")

	return router

}
