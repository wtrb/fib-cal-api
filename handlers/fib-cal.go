package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	db "github.com/wtrb/fib-cal-api/database"
)

func FibCal(c *gin.Context) {
	idx := c.Param("index")

	if n, err := strconv.Atoi(idx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else if n >= 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "index must be lower than 100",
		})
		return
	}

	db.InsertValue(idx)
	db.Queue(idx)
}
