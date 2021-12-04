package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/wtrb/fib-cal-api/routes"
)

const (
	srvName = "fib-cal-api"
)

func main() {
	log.Println(srvName + " starting")

	r := gin.Default()
	routes.SetRoutes(r)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
