// routers/router.go

package routers

import (
	"GoBoard/controllers"
	auth "GoBoard/middlewares"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	router.Use(auth.SetUserStatus())

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.

	// Handle the index route
	router.GET("/", controllers.ShowIndexPage)

	articleRoutes := router.Group("/article")
	{
		// route from Part 1 of the tutorial
		articleRoutes.GET("/view/:article_id", controllers.GetArticle)

		articleRoutes.GET("/create", controllers.ShowArticleCreationPage)

		articleRoutes.POST("/create", controllers.CreateArticle)
	}

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", controllers.ShowRegistrationPage)
		userRoutes.POST("/register", controllers.Register)

		userRoutes.GET("/login", controllers.ShowLoginPage)
		userRoutes.POST("/login", controllers.PerformLogin)
		userRoutes.GET("/logout", controllers.Logout)
	}

	return router
}
