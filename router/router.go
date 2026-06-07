package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	// Initialize Router for Gin
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Running the Server
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	router.Run(":8080")
}
