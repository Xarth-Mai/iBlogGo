package database

import (
	"github.com/Xarth-Mai/iBlogGo/models"
	"gorm.io/gorm"
)

// GetPostList 返回文章列表数据
func GetPostList(db *gorm.DB) (models.List, error) {
	var list models.List
	// 查找所有文章的标题和日期
	err := db.Select("id", "title", "date").Order("date DESC").Limit(20).Find(&list).Error
	return list, err
}
