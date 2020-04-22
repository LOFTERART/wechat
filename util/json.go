package util

import (
	"encoding/json"
)

// 生成JSON字符串
func MarshalJson(m interface{}) string {
	str, _ := json.Marshal(m)
	return string(str)
}
