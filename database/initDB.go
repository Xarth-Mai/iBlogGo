package database

import (
	"github.com/Xarth-Mai/iBlogGo/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"time"
)

// InitDB 初始化数据库并返回一个指向数据库连接的指针
func InitDB() (*gorm.DB, error) {
	// 使用 SQLite 数据库
	db, err := gorm.Open(sqlite.Open("iBlogGo.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移模型，确保数据库表格与结构体一致
	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		return nil, err
	}

	// 检查数据库数据量，如果为空则插入测试数据
	var count int64
	db.Model(&models.Post{}).Count(&count)
	println("Number of posts in the database:", count)
	if count == 0 {
		println("Database is empty. Inserting test data...")
		return InsertTestPost(db)
	}

	return db, nil
}

// InsertTestPost 插入测试数据
func InsertTestPost(db *gorm.DB) (*gorm.DB, error) {
	err := db.Create(&models.Post{
		Title:   "First Post",
		Content: "This is the first post.",
		Date:    time.Now(),
	}).Error
	if err != nil {
		return nil, err
	}

	err = db.Create(&models.Post{
		Title:   "Second Post",
		Content: "This is the Second post.",
		Date:    time.Date(2050, 10, 10, 0, 0, 0, 0, time.UTC),
	}).Error
	if err != nil {
		return nil, err
	}
	return db, nil
}
