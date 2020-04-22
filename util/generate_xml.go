package util

import (
	"bytes"
	"fmt"
)

// 生成请求XML的Body体
func GenerateXml(data map[string]interface{}) string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml>")
	for k, v := range data {
		buffer.WriteString(fmt.Sprintf("<%s><![CDATA[%v]]></%s>", k, v, k))
	}
	buffer.WriteString("</xml>")
	return buffer.String()
}
