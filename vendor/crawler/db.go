package crawler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

// 模块表结构
type Modules struct {
	Id    uint
	Name  string
	Alias string
}

var DB *gorm.DB

func init() {
	DB = db()
}

func db() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/home/vsuper/work/go/src/stockcrawler/crawler.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return db
}
