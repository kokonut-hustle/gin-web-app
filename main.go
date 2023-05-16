package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded from the disk again.
	// This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// To start with, we'll use an inline route handler.
	// Later on, we'll create standalone functions that will be used as route handlers.
	router.GET("/", showIndexPage)

	router.Run()
}

func showIndexPage(c *gin.Context) {
	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Home Page",
		},
	)
}
