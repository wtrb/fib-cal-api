package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/wtrb/fib-cal-api/handlers"
)

func SetRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/healthz", handlers.GetHealthz)

	api := r.Group("/api")
	{
		api.GET("/values", handlers.GetAllValues)
		api.GET("/values/:index", handlers.GetValueByIndex)
		api.POST("/cal/:index", handlers.FibCal)
	}
}
