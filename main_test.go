package main

import (
	"crawlers"
	// "share/crawlers"
	"testing"
)

func Test_reportIndustry(t *testing.T) {
	
	module := &crawlers.ReportIndustry{}
	module.Run()
}
