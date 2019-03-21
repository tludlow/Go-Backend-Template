package router

import (
	"billy-maths/database"
	"log"
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
		v1.POST("/enter/:name", enterNewUser)
	}
}

func enterNewUser(c *gin.Context) {
	name := c.Param("name")

	statement, err := database.Database.Prepare("INSERT INTO users (name) VALUES (?)")

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Failed to prepare to add the user '" + name + "' to the database",
		})
		return
	}

	_, err = statement.Exec(name)

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Failed to add the user '" + name + "' to the database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Inserted the new user into the database",
	})
	return
}
