package wechat

import (
	"bytes"
	"fmt"
	"sort"
)

// 本地通过支付参数计算签名值
func (c *PayClient) localSign(body map[string]interface{}, signType string, apiKey string) string {
	signStr := c.sortSignParams(body, apiKey)
	return SignWithType(signType, signStr, apiKey)
}

// 获取根据Key排序后的请求参数字符串
func (c *PayClient) sortSignParams(body map[string]interface{}, apiKey string) string {
	keyList := make([]string, 0)
	for k := range body {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		s := fmt.Sprintf("%s=%s&", k, fmt.Sprintf("%v", body[k]))
		buffer.WriteString(s)
	}
	buffer.WriteString(fmt.Sprintf("key=%s", apiKey))
	return buffer.String()
}
