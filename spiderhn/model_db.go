package spiderhn

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var db *gorm.DB

func LoadSQLiteDb(verbose bool) {
	log.Println("SQLite3 in:", "sqlite3.db")
	sqlite, err := gorm.Open("sqlite3", "sqlite3.db")
	if err != nil {
		log.Fatal(err)
		return
	}

	db = sqlite
	//TODO::optimize
	//db.DropTable("term_logs")
	db.AutoMigrate(HackNew{})
	db.LogMode(verbose)
}

type BaseModel struct {
	Id        uint      `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
