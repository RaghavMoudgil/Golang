package routes

import (
	"mongoDB/controller"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	router.GET("/api/movies", controller.GetAllMoviesController)
	router.POST("/api/movies", controller.CreateMovie)
	router.PUT("/api/movies/:id", controller.MarkAsWatched)
	router.DELETE("api/movies/:id", controller.DeleteOneMovie)
	router.DELETE("api/deleteallmovie", controller.DeleteAllMovies)
	router.Run(":8082")
}
