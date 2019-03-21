package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
* This file will house all of the API routes we will be using. The API will be versioned with a prefix such as "DOMAIN/v1/ROUTE" so that we can
* ensure API availability, if we have new routes being created for a large incremental version of the application then we will increase the API versioning.
* The end goal will be creating a micro-service architecture for the backend, keep this in mind when developing the code for the routes and backend in general.
 */

func setupRoutesV1() {
	v1 := Router.Group("/v1")
	{
		v1.GET("/testing/:name", testRoute)
	}
}

func testRoute(c *gin.Context) {
	givenName := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, and welcome to Billy-Maths",
		"name":    "Your name is, " + givenName,
	})
}
