package database

import (
	"github.com/Xarth-Mai/iBlogGo/models"
	"gorm.io/gorm"
)

// GetPost 返回指定 ID 的文章数据
func GetPost(db *gorm.DB, id string) (models.Post, error) {
	var post models.Post
	// 查找指定 ID 的文章
	err := db.Find(&post, id).Error
	return post, err
}
