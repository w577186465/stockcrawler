package model

import (
	"time"
)

type ReportIndustry struct {
	Id          uint
	Title       string
	Pjchange    string
	Insname     string // 券商名称
	Indid       string // 行业id
	Pjtype      string
	Expect      string
	Indname     string // 行业名称
	Fluctuation string // 涨跌幅
	Hash        string // 涨跌幅
	Content     string
	CreatedAt   time.Time
}
