package wechat

import (
	"encoding/json"
	"time"
)

// 生成模型字符串
func JsonString(m interface{}) string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}

// 格式化时间，按照yyyyMMddHHmmss格式
func FormatDateTime(t time.Time) string {
	return t.Format("20060102150405")
}
