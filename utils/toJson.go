package utils

import (
	"encoding/json"
	"log"
)

// ToJson 将结构体或结构体数组转换为 JSON 字符串
func ToJson(a interface{}) string {
	jsonData, err := json.Marshal(a)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return "" // 返回空字符串
	}
	return string(jsonData) // 正常返回 JSON 字符串
}
