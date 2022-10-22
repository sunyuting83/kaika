package controller

import (
	LevelDB "kaika/Leveldb"
	"kaika/database"
	"kaika/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Node node
type Login struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func Sgin(c *gin.Context) {
	var form Login
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	if len(form.UserName) <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't username",
		})
		return
	}
	if len(form.Password) <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't password",
		})
		return
	}
	user, err := LevelDB.Get(form.UserName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "登陆失败",
		})
		return
	}
	if user != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "登陆成功",
			"token":   user,
		})
		return
	}
	secret_key, _ := c.Get("SECRET_KEY")
	SECRET_KEY := secret_key.(string)
	PASSWD := utils.MD5(strings.Join([]string{form.Password, SECRET_KEY}, ""))
	var admin database.Admin
	login, err := admin.CheckAdminLogin(form.UserName, PASSWD)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "登陆失败",
		})
		return
	}
	T := time.Now().Format("202210100000")
	token := strings.Join([]string{login.Username, login.Password, T}, "")
	token = utils.MD5(token)
	var ttl int64 = 1000 * 60 * 60 * 24 * 90

	// ASE加密token
	TOKEN, err := utils.AesEncrypt([]byte(token), []byte(SECRET_KEY))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "登陆失败",
		})
		return
	}

	LevelDB.Set(form.UserName, token, ttl)
	LevelDB.Set(token, token, ttl)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "登陆成功",
		"token":   string(TOKEN),
	})
}
