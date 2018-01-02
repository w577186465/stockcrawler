package main

import (
	"crawler"
	"crawlers"
	"testing"
)

type TestModel struct {
	ID      uint
	Content string
}

func Test_reportIndustry(t *testing.T) {
	module := &crawlers.ReportIndustry{}
	module.Run()
}

func Test_createLink(t *testing.T) {
	crawler.CreateModule(true, "testb", &TestModel{})
}

func Test_createHash(t *testing.T) {
	crawler.CreateModule(false, "test", &TestModel{})
}