// router.router.go

package router

import (
	"GoBoard/src/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.

	// Handle the index route
	router.GET("/", handler.ShowIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", handler.GetArticle)

	return router
}
