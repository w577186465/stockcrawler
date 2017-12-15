package crawler

import (
	"fmt"
)

type Module struct {
	Name string
}

// 创建模块 参数：name模块名称，model信息库Model
func CreateModule(name string, model interface{}) {
	dataTable := fmt.Sprintf("%s_data", name)  // 信息库表名
	linkTable := fmt.Sprintf("%s_links", name) // 链接库表名

	// 创建信息库
	if err := DB.Table(dataTable).CreateTable(model).Error; err != nil {
		fmt.Println(err)
	}

	CreateLinkTable(linkTable) // 创建链接库

	fmt.Printf("%s模块创建成功\r\n", name)
}

// 创建链接库
func CreateLinkTable(name string) {
	DB.Table(name).CreateTable(&Link{})
}

// 增加链接
func (m Module) Addlink(link, hash string) {
	name := m.Name
	table := fmt.Sprintf("%s_links", name)

	// 判断重复
	find := Link{
		Hash: hash,
	}
	DB.Table(table).Find(&find)

	add := Link{Link: link, Hash: hash}
	DB.Table(table).Create(&add)
}

// 设置链接采集成功状态
func (m Module) LinkSuccess(id int) {
	table := fmt.Sprintf("%s_links", m.Name)
	update := Link{State: 1}
	DB.Table(table).Where("id = ?", id).Update(&update)
}

func (m Module) AddData(data interface{}) {
	table := fmt.Sprintf("%s_data", m.Name)
	fmt.Println(DB.Table(table).Create(data).Error)
}
