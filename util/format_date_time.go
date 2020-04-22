package util

import (
	"time"
)

// 格式化时间，按照yyyyMMddHHmmss格式
func FormatDateTime(t time.Time) string {
	return t.Format("20060102150405")
}
