package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"flag"
)

var DB *gorm.DB

// 模块表结构
type Modules struct {
	Id    uint
	Name  string
	Alias string
}

func init () {
	ConnectDB("crawler.db")
}

func main() {
	var init bool
	flag.BoolVar(&init, "init", false, "初始化数据库")
	flag.Parse()
	fmt.Println(init)
	if init {
		intdb()
	}

	var module string
	flag.StringVar(&module, "make:module", "", "创建模块")
	fmt.Println(module)
	if module != "" {
		// createModule()
	}
}

func ConnectDB(f string) *gorm.DB {
	var err error
	DB, err = gorm.Open("sqlite3", f)
	if err != nil {
		panic(err)
	}
	return DB
}

// 初始化数据库
func intdb() {
	if DB.HasTable("modules") {
		fmt.Println("操作执行过了")
		return
	}
	DB.CreateTable(&Modules{}) // 创建模块
	fmt.Println("初始化成功")
}

// // 创面模块
// func createModule(name, alias string, model interface{}) {
// 	// 判断重复
// 	find := Modules{
// 		Name: name,
// 	}
// 	DB.Table(table).Find(&find)
// 	fmt.Println(find)

// 	DB.Create(&Modules{Name: name, Alias: alias}) // 添加模块
// 	dataTable := fmt.Sprintf("%s_data", name)  // 信息库表名
// 	linkTable := fmt.Sprintf("%s_links", name) // 链接库表名

// 	// 创建信息库
// 	if err := DB.Table(dataTable).CreateTable(model).Error; err != nil {
// 		fmt.Println(err)
// 	}

// 	CreateLinkTable(linkTable) // 创建链接库

// 	fmt.Printf("%s模块创建成功\r\n", name)
// }

// // 创建链接库
// func CreateLinkTable(name string) {
// 	DB.Table(name).CreateTable(&Link{})
// }