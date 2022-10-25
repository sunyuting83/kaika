package router

import (
	controller "kaika/controller"
	Admin "kaika/controller/Admin"
	Card "kaika/controller/Card"
	utils "kaika/utils"

	"github.com/gin-gonic/gin"
)

// SetConfigMiddleWare set config
func SetConfigMiddleWare(SECRET_KEY string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("secret_key", SECRET_KEY)
		c.Writer.Status()
	}
}

// InitRouter make router
func InitRouter(SECRET_KEY string) *gin.Engine {
	router := gin.Default()
	router.Use(utils.CORSMiddleware())
	api := router.Group("/api")
	api.Use(SetConfigMiddleWare(SECRET_KEY))
	{
		router.GET("/", controller.Index)
		api.POST("/addadmin", utils.VerifyMiddleware(), Admin.AddAdmin)
		api.PUT("/repassword", utils.VerifyMiddleware(), Admin.ResetPassword)
		api.DELETE("/deladmin", utils.VerifyMiddleware(), Admin.DeleteAdmin)
		api.POST("/loginadmin", Admin.Sgin)
		api.POST("/creatdcard", utils.VerifyMiddleware(), Card.CreateCard)
		api.GET("/cardlist", utils.VerifyMiddleware(), Card.CardList)
		api.GET("/checklogin", utils.VerifyMiddleware(), Admin.CheckLogin)
		api.GET("/login", Card.Login)
		api.GET("/check", Card.CheckCard)
	}

	return router
}
