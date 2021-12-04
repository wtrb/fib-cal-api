package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/wtrb/fib-cal-api/database"
)

func GetValueByIndex(c *gin.Context) {
	val := db.SelectValueByIndex(c.Param("index"))
	c.JSON(http.StatusOK, gin.H{
		"value": val,
	})
}
