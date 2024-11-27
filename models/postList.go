package models

type List []Posts

type Posts struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}
