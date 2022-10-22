package router

import (
	controller "kaika/controller"
	Admin "kaika/controller/Admin"
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
		// api.DELETE("/delone", utils.VerifyMiddleware(), controller.DeleteOne)
		api.POST("/login", Admin.Sgin)
	}

	return router
}
