package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// PageHandler 处理根路径和 /blog/* 路径的请求
func PageHandler(w http.ResponseWriter, r *http.Request) {
	// 只接受处理 / 和 /blog/* 路径
	isBlog := strings.HasPrefix(r.URL.Path, "/blog/")
	if r.URL.Path != "/" && !isBlog {
		http.NotFound(w, r)
		log.Printf("Invalid path: %s", r.URL.Path)
		return
	}

	// 设置响应内容类型为 text/html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 根据路径动态设置模板文件
	var filenames string
	if isBlog {
		// 对于 /blog/* 请求，可以加载不同的模板（此处示例使用相同模板）
		filenames = "views/blog.html"
	} else {
		// 根路径请求加载 index.html
		filenames = "views/index.html"
	}

	// 解析模板文件
	index, err := template.ParseFiles(filenames)
	if err != nil {
		log.Printf("Failed to parse template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 执行模板并写入响应
	if err := index.Execute(w, nil); err != nil {
		log.Printf("Failed to execute template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
