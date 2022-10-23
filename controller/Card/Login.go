package controller

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	var a string = c.Query("a")
	var b string = c.Query("b")
	c.String(200, a+"失败"+b)
}
