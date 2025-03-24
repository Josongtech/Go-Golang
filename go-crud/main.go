package main

import (
	"example.com/m/controllers"
	"example.com/m/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.GET("/", controllers.PostsCreate())

	r.Run()
}
