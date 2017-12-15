package main

import (
	"share/crawlers"
	// "share/crawlers"
	"testing"
)

func Test_reportIndustry(t *testing.T) {
	module := &crawler.ReportIndustry{}
	module.Run()
}
