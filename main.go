// main.go

package main

import (
	"GoBoard/configs"
	"GoBoard/database"
	"GoBoard/routers"
)

func main() {
	configs.Init()
	database.Init()

	// Initialize the routes
	var router = routers.InitializeRoutes()
	// Start serving the application
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}
