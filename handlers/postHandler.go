package handlers

import (
	"github.com/Xarth-Mai/iBlogGo/database"
	"github.com/Xarth-Mai/iBlogGo/utils"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// PostHandler 处理文章HTTP请求
func PostHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id := r.URL.Path
	id = id[len("/blog/"):]
	if id == "" {
		http.NotFound(w, r)
		return
	}

	// 从数据库中获取文章
	post, err := database.GetPost(db, id)
	if err != nil {
		log.Printf("Failed to get post: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 解析Markdown格式文章
	post.Content = utils.ParseMarkdown(post.Content)

	// 将文章转换为JSON格式并写入响应
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(utils.ToJson(&post)))
	if err != nil {
		log.Printf("Failed to write response: %v", err)
		return
	}

}
