package models

type Post struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}
