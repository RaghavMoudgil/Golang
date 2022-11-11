package main

import (
	"mongoDB/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Router(router)
	router.Run("localhost:8082")

}
