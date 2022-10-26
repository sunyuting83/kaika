package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTimeDate(c *gin.Context) {
	T := time.Now().Unix()
	c.JSON(http.StatusOK, gin.H{
		"datetime": T,
	})
}
