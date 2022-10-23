package controller

import (
	LevelDB "kaika/Leveldb"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckCard(c *gin.Context) {
	var a string = c.Query("a")
	var b string = c.Query("b")
	if len(a) != 32 && len(b) != 32 {
		c.String(200, "1")
		return
	}
	data, err := LevelDB.CardGet(a, b, false)
	if err != nil {
		c.String(200, "1")
		return
	}
	if len(string(data)) <= 0 {
		c.String(200, "1")
		return
	}
	comput := strings.Split(string(data), "----")[0]
	if comput != b {
		c.String(200, "1")
		return
	}
	c.String(200, "0")
}
