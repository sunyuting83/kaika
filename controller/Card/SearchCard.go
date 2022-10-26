package controller

import (
	"kaika/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Search struct {
	Card string `form:"card" json:"card" xml:"card"  binding:"required"`
}

func SearchCard(c *gin.Context) {
	var form Search
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	if len(form.Card) != 32 {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "非法卡号",
		})
		return
	}
	var card database.Card
	data, err := card.Search(form.Card)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "失败",
		})
		return
	}
	Data := gin.H{
		"status": 0,
		"data":   data.Card,
	}
	c.JSON(http.StatusOK, Data)
}
