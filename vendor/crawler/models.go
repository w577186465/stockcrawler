package crawler

import (
	// "time"
)

// 链接库
type Link struct {
	ID        uint
	Link      string
	Hash      string `gorm:"type:char(64);index;unique"`
	State     int
}

type Hash struct {
	ID        uint
	Hash      string `gorm:"type:char(64);index;unique"`
}