package main

import (
	"github.com/Xarth-Mai/iBlogGo/handlers"
	"log"
	"net/http"
)

func main() {
	// HTML模板
	http.HandleFunc("/", handlers.PageHandler)

	// 静态资源
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// 文章列表JSON
	http.HandleFunc("/list", handlers.PostListHandler)

	// 博客文章JSON
	http.HandleFunc("/post/", handlers.PostHandler)

	// 启动服务器
	println("服务器启动在 http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
