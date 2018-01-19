package main

import (
	"crawler"
	"crawlers"
	"fmt"
	// "net/http"
	"testing"
)

type TestModel struct {
	ID      uint
	Content string
}

func Test_initdb(t *testing.T) {
	crawler.Initdb()
}

func Test_http(t *testing.T) {
	var i = 0
	for {
		i++
		fmt.Println(i)
		url := "http://localhost/"
		crawler.Request(url).Retry(10).Delay(10).Download().Json()
	}
}

func Test_reportIndustry(t *testing.T) {
	module := &crawlers.ReportIndustry{
		Name: "report_industry",
	}
	module.Run()
}

func Test_createLink(t *testing.T) {
	crawler.CreateModule(true, "testb", "测试b", &TestModel{})
}

func Test_createHash(t *testing.T) {
	crawler.CreateModule(false, "test", "测试", &TestModel{})
}

func Test_getModule(t *testing.T) {
	module := crawler.GetModule("report_industry")
	fmt.Println(module)
}
