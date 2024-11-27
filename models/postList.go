package models

import "time"

type List []PostMetadata

type PostMetadata struct {
	ID    int       `gorm:"primaryKey" json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

// TableName 方法用于指定表名
func (PostMetadata) TableName() string {
	return "posts" // 显式指定查询的表名为 "posts"
}
