// main.go

package main

import (
	"GoBoard/src/router"
)

func main() {

	// Initialize the routes
	var router = router.InitializeRoutes()

	// Start serving the application
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}
