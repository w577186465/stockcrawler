package crawler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var DB *gorm.DB

// 模块表结构
type Modules struct {
	Id    uint
	Name  string
	Alias string
}

func init () {
	ConnectDB("/home/vsuper/work/go/src/stockcrawler/crawler.db")
}

func ConnectDB(f string) *gorm.DB {
	var err error
	DB, err = gorm.Open("sqlite3", f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return DB
}
