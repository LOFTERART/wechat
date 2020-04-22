package util

import (
	"encoding/hex"
	"gitee.com/xiaochengtech/wechat/constant"
	"strings"
)

// 根据签名类型，生成签名
func SignWithType(signType string, origin string, apiKey string) string {
	var hashSign []byte
	if signType == constant.SignTypeHmacSHA256 {
		hashSign = HmacSha256(origin, apiKey)
	} else {
		hashSign = Md5(origin)
	}
	return strings.ToUpper(hex.EncodeToString(hashSign))
}
