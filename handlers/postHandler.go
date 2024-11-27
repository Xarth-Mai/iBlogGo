package handlers

import (
	"github.com/Xarth-Mai/iBlogGo/database"
	"github.com/Xarth-Mai/iBlogGo/utils"
	"log"
	"net/http"
)

// PostHandler 获取文章
func PostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	id = id[len("/blog/"):]
	if id == "" {
		http.NotFound(w, r)
		return
	}

	// 从数据库中获取文章
	post, err := database.GetTestPost(id)
	if err != nil {
		log.Printf("Failed to get post: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 将文章转换为JSON格式并写入响应
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(utils.ToJson(&post)))
	if err != nil {
		log.Printf("Failed to write response: %v", err)
		return
	}

}
