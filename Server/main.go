package main

import (
	"assessment-bootcamp/Server/handler"
	"assessment-bootcamp/Server/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(handler.CorsMiddleware())

	routes.UserRoute(r)
	routes.PasswordRoute(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
