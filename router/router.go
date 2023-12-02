package router

import (
	"github.com/gin-gonic/gin"
)

func Inicialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize routes
	initialize(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
