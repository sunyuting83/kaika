package main

import (
	"fmt"
	LevelDB "kaika/Leveldb"
	orm "kaika/database"
	"kaika/router"
	"kaika/utils"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	OS := runtime.GOOS
	CurrentPath, _ := utils.GetCurrentPath()

	confYaml, err := utils.CheckConfig(OS, CurrentPath)
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Duration(10) * time.Second)
		os.Exit(0)
	}
	pwd := utils.MD5(strings.Join([]string{confYaml.AdminPWD, confYaml.SECRET_KEY}, ""))
	orm.InitDB(CurrentPath, pwd)
	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)
	defer orm.Eloquent.Close()
	defer LevelDB.LevelDB.Close()
	app := router.InitRouter(confYaml.SECRET_KEY)

	app.Run(strings.Join([]string{":", confYaml.Port}, ""))
}
