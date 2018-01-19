package crawler

import (
	"fmt"
)

type Module struct {
	Id        int
	Name      string
	Alias     string
	HashTable string
	DataTable string
}

// 初始化数据库
func Initdb() {
	if DB.HasTable("modules") {
		fmt.Println("操作执行过了")
		return
	}
	DB.CreateTable(&Module{}) // 创建模块
	fmt.Println("初始化成功")
}

// 获取模块信息
func GetModule(name string) Module {
	module := Module{}
	DB.Where("name = ?", name).Find(&module)
	return module
}

// 创建模块 参数：name模块名称，model信息库Model
func CreateModule(link bool, name string, alias string, model interface{}) {
	dataTable := fmt.Sprintf("%s_data", name)  // 信息库表名
	linkTable := fmt.Sprintf("%s_hashs", name) // 链接库表名

	module := Module{Name: name, Alias: alias, HashTable: linkTable, DataTable: dataTable}
	if err := DB.Create(&module).Error; err != nil {
		panic(err)
	}

	// 创建信息库
	if err := DB.Table(dataTable).CreateTable(model).Error; err != nil {
		panic(err)
	}

	if link {
		CreateLinkTable(linkTable) // 创建链接库
	} else {
		CreateHashTable(linkTable) // 创建链接库
	}

	fmt.Printf(" %s 模块创建成功\r\n", alias)
}

// 创建链接库
func CreateLinkTable(name string) {
	if err := DB.Table(name).CreateTable(&Link{}).Error; err != nil {
		panic(err)
	}
}

// 创建哈希库
func CreateHashTable(name string) {
	if err := DB.Table(name).CreateTable(&Hash{}).Error; err != nil {
		panic(err)
	}
}

// 增加链接
func (m Module) Addlink(link, hash string) {
	table := m.HashTable

	// 判断重复
	find := Link{
		Hash: hash,
	}
	DB.Table(table).Find(&find)

	add := Link{Link: link, Hash: hash}
	DB.Table(table).Create(&add)
}

// 增加哈希值
func (m Module) AddHash(hash string) bool {
	table := m.HashTable

	// 判断重复
	find := Link{}
	DB.Table(table).Where("hash = ?", hash).First(&find)
	if find.ID != 0 {
		return false
	}

	add := Hash{Hash: hash}
	DB.Table(table).Create(&add)
	return true
}

func (m Module) HasHash(hash string) bool {
	table := m.HashTable
	find := Link{}
	DB.Table(table).Where("hash = ?", hash).First(&find)
	if find.ID != 0 {
		return false
	}
	return true
}

// 设置链接采集成功状态
func (m Module) LinkSuccess(id int) {
	table := m.HashTable // 哈希库
	update := Link{State: 1}
	DB.Table(table).Where("id = ?", id).Update(&update)
}

func (m Module) AddData(data interface{}) {
	table := m.DataTable // 数据表
	DB.Table(table).Create(data)
}
