package controller

import (
	"kaika/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpStatusAdmin(c *gin.Context) {
	var form User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	var (
		admin     database.Admin
		Status    int    = 1
		StatusStr string = "锁定"
	)
	user, err := admin.CheckID(form.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	if user.Status == 1 {
		Status = 0
		StatusStr = "解锁"
	}
	admin.Status = Status
	a, err := admin.UpStatusOne(user.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
			"user":    a,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": strings.Join([]string{"成功", StatusStr, "管理员"}, ""),
		"user":    a,
	})
}
