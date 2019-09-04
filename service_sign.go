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

// JSAPI支付，统一下单获取支付参数后，再次计算出小程序用的paySign
func GetMiniPaySign(appId, nonceStr, prepayId, signType, timeStamp, apiKey string) (paySign string) {
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

// JSAPI支付，统一下单获取支付参数后，再次计算出微信内H5支付需要用的paySign
func GetH5PaySign(appId, nonceStr, packages, signType, timeStamp, apiKey string) (paySign string) {
	// 原始字符串
	raw := fmt.Sprintf("appId=%s&nonceStr=%s&package=%s&signType=%s&timeStamp=%s&key=%s",
		appId, nonceStr, packages, signType, timeStamp, apiKey)
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

// APP支付，统一下单获取支付参数后，再次计算APP支付所需要的的sign
func GetAppPaySign(appId, nonceStr, partnerId, prepayId, signType, timeStamp, apiKey string) (paySign string) {
	// 原始字符串
	raw := fmt.Sprintf("appId=%s&nonceStr=%s&package==Sign=WXPay&partnerid=%s&prepayid=%s&timeStamp=%s&key=%s",
		appId, nonceStr, partnerId, prepayId, timeStamp, apiKey)
	buffer := new(bytes.Buffer)
	buffer.WriteString(raw)
	// 加密签名
	signStr := buffer.String()
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
