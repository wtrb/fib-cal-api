package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/wtrb/fib-cal-api/database"
)

func GetAllValues(c *gin.Context) {
	vals := db.SelectAllValues()
	c.JSON(http.StatusOK, gin.H{
		"values": vals,
	})
}
