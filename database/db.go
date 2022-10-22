package database

import (
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	Eloquent *gorm.DB
)

// InitDB init db
func InitDB(d string) {
	dbName := strings.Join([]string{d, "data/data.sqlite"}, "/")
	// fmt.Println(dbName)
	Eloquent, _ = gorm.Open("sqlite3", dbName)
	if !Eloquent.HasTable(&Admin{}) && !Eloquent.HasTable(&Card{}) {
		newTime := time.Now().Unix()
		var admin Admin
		admin.Username = "sleepsun"
		admin.Password = "b346ed32af6bb2b4023f0fbd32866a81"
		admin.Status = 0
		admin.UpdateTime = newTime
		admin.CreatedTime = newTime

		if err := Eloquent.CreateTable(&Admin{}).Error; err != nil {
			log.Fatal(err)
		}
		if err := Eloquent.CreateTable(&Card{}).Error; err != nil {
			log.Fatal(err)
		}
		admin.Insert()
	}
	Eloquent.SingularTable(true)
}
