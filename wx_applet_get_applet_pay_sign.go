package wechat

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// 小程序支付，统一下单获取支付参数后，再次计算出小程序用的paySign
func GetAppletPaySign(appId, nonceStr, prepayId, signType, timeStamp, apiKey string) (paySign string) {
	// 原始字符串
	raw := fmt.Sprintf("appId=%s&nonceStr=%s&package=%s&signType=%s&timeStamp=%s&key=%s",
		appId, nonceStr, prepayId, signType, timeStamp, apiKey)
	buffer := new(bytes.Buffer)
	buffer.WriteString(raw)
	signStr := buffer.String()
	// 加密签名
	var hashSign []byte
	if signType == SignTypeHmacSHA256 {
		hash := hmac.New(sha256.New, []byte(apiKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	paySign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}
