package controller

import (
	LevelDB "kaika/Leveldb"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var a string = c.Query("a")
	var b string = c.Query("b")
	if len(a) != 32 && len(b) != 32 {
		c.String(200, "1")
		return
	}
	data, err := LevelDB.CardGet(a, b, true)
	if err != nil {
		c.String(200, "1")
		return
	}
	if len(string(data)) <= 0 {
		c.String(200, "1")
		return
	}
	c.String(200, "0")
}
