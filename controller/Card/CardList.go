package controller

import (
	"kaika/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CardList(c *gin.Context) {
	var page string = c.DefaultQuery("page", "0")
	pageInt, _ := strconv.ParseInt(page, 10, 64)
	var card database.Card
	count, err := card.GetCount()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "失败",
		})
		return
	}
	dataList, err := card.GetCardList(pageInt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "失败",
		})
		return
	}
	Data := gin.H{
		"status": 1,
		"data":   dataList,
		"total":  count,
	}
	c.JSON(http.StatusOK, Data)
}
