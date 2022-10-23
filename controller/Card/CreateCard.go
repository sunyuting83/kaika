package controller

import (
	LevelDB "kaika/Leveldb"
	"kaika/database"
	"kaika/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Node node
type Card struct {
	Card     int   `form:"card" json:"card" xml:"card"  binding:"required"`
	DateTime int64 `form:"datetime" json:"datetime" xml:"datetime"  binding:"required"`
}

func CreateCard(c *gin.Context) {
	var form Card
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	if form.Card < 1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't card",
		})
		return
	}
	if form.DateTime < 86400 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "不能小于1天",
		})
		return
	}
	secret_key, _ := c.Get("secret_key")
	SECRET_KEY := secret_key.(string)
	var (
		card *database.Card
	)
	for i := 0; i < form.Card; i++ {
		T := time.Now().Unix()
		TimeStr := strconv.FormatInt(T, 10)
		rand := utils.RandSeq(16)
		Card := utils.MD5(strings.Join([]string{TimeStr, SECRET_KEY, rand}, ""))
		Card = strings.ToUpper(Card)
		card = &database.Card{Card: Card, UpdateTime: T + form.DateTime, CreatedTime: T}
		id, _ := card.Insert()
		ID := strconv.FormatInt(id, 10)
		value := strings.Join([]string{"nil", ID}, "----")
		LevelDB.Set(Card, value, form.DateTime)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "开卡成功",
	})
}
