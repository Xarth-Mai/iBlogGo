package main

import (
	"fmt"
	"github.com/Xarth-Mai/iBlogGo/database"
	"github.com/Xarth-Mai/iBlogGo/handlers"
	"log"
	"net/http"
)

func main() {

	port := "6102"

	// 初始化数据库连接
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// HTML模板
	http.HandleFunc("/", handlers.PageHandler)

	// 静态资源
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// 文章列表JSON
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostListHandler(w, r, db)
	})

	// 博客文章JSON
	http.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostHandler(w, r, db)
	})

	// 启动服务器
	addr := fmt.Sprintf("127.0.0.1:%s", port)
	println("Server started at http://" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
