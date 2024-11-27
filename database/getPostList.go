package database

import (
	"github.com/Xarth-Mai/iBlogGo/models"
	"gorm.io/gorm"
)

// GetPostList 返回文章列表数据
func GetPostList(db *gorm.DB) (models.List, error) {
	var posts models.List
	// 查找所有文章的标题和日期
	if err := db.Select("id", "title", "date").Order("date DESC").Limit(20).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
