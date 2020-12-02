// main.go

package main

import (
	"GoBoard/routers"
)

func main() {

	// Initialize the routes
	var router = routers.InitializeRoutes()

	// Start serving the application
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}
