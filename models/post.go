package models

import "time"

type Post struct {
	ID      int       `gorm:"primaryKey" json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}
