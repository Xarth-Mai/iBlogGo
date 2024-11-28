package utils

import (
	"github.com/russross/blackfriday/v2"
)

// ParseMarkdown 将 Markdown 解析为 HTML
func ParseMarkdown(content string) string {
	htmlContent := blackfriday.Run([]byte(content))
	return string(htmlContent)
}
