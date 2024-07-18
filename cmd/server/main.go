package main

import (
	"log"
	"partner-locator/internal/app"
	"partner-locator/internal/infrastructure/router"

	"github.com/gin-gonic/gin"
)

func main() {
	application := app.NewApplication()

	routerGin := gin.Default()

	// Initialize the routes
	router.InitRoutes(routerGin, application)

	log.Println("Starting server on :8080")
	log.Fatal(routerGin.Run(":8080"))
}
