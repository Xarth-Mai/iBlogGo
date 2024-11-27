package handlers

import (
	"github.com/Xarth-Mai/iBlogGo/database"
	"github.com/Xarth-Mai/iBlogGo/utils"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// PostListHandler 处理文章列表的HTTP请求
func PostListHandler(w http.ResponseWriter, _ *http.Request, db *gorm.DB) {
	// 从数据库中获取文章列表
	list, err := database.GetPostList(db)
	if err != nil {
		log.Printf("Failed to get list: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 将列表转换为JSON格式并写入响应
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(utils.ToJson(&list)))
	if err != nil {
		log.Printf("Failed to write response: %v", err)
		return
	}
}
