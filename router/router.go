package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

//Router - Our exposed gin router struct which can be used across multiple files.
var Router *gin.Engine

//CreateRouter - Establish a new HTTP router using the gin framework for our api.
func CreateRouter() {
	Router = gin.Default()

	//We have a seperate file for storing the API routes (router/routes.go), we will need to make the router aware of them next
	//Before we start the router, lets do that.
	setupRoutesV1()

	//We have our required port stored within the .env file, load that as a string so the router can use it.
	httpPort := os.Getenv("HTTP_PORT")
	Router.Run(":" + httpPort)
	log.Println("Starting the HTTP Router.")
}
