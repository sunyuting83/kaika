package database

import (
	"log"
	"strings"

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
		if err := Eloquent.CreateTable(&Admin{}).Error; err != nil {
			log.Fatal(err)
		}
		if err := Eloquent.CreateTable(&Card{}).Error; err != nil {
			log.Fatal(err)
		}
	}
	Eloquent.SingularTable(true)
}
