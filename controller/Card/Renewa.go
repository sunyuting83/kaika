package controller

import (
	LevelDB "kaika/Leveldb"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Cards struct {
	Card     string `form:"card" json:"card" xml:"card"  binding:"required"`
	DateTime int64  `form:"datetime" json:"datetime" xml:"datetime"  binding:"required"`
}

func RenewaCard(c *gin.Context) {
	var form Cards
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	if len(form.Card) != 32 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "卡号错误",
		})
		return
	}
	// 86400
	if form.DateTime < 86400 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "不能小于1天",
		})
		return
	}
	data, err := LevelDB.Get(form.Card)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "卡号已过期或不存在",
		})
		return
	}
	if data == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "卡号已过期或不存在",
		})
		return
	}
	d, err := LevelDB.RenewaGet(form.Card, form.DateTime)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "卡号已过期或不存在",
		})
		return
	}
	if d == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "卡号已过期或不存在",
		})
		return
	}
	NewData := strings.Split(string(d), "----")
	NewID, _ := strconv.ParseInt(NewData[0], 10, 64)
	NewTime, _ := strconv.ParseInt(NewData[1], 10, 64)
	c.JSON(http.StatusOK, gin.H{
		"status":     0,
		"message":    "续费成功",
		"id":         NewID,
		"updatetime": NewTime,
	})
}
