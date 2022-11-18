package main

import (
	"mongoDB/middleware"
	"mongoDB/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	routes.Router(router)
	router.Run("localhost:8082")

}
