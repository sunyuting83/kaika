package controller

import (
	"kaika/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
}

func DeleteAdmin(c *gin.Context) {
	var form User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	if len(form.UserName) <= 4 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't username",
		})
		return
	}
	var admin database.Admin
	user, err := admin.CheckUserName(form.UserName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	admin.DeleteOne(form.UserName)
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":   0,
		"message":  "成功删除管理员",
		"username": user.Username,
	})
}
