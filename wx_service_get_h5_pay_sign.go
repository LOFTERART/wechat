package wechat

import (
	"bytes"
	"fmt"
)

// JSAPI支付，统一下单获取支付参数后，再次计算出微信内H5支付需要用的paySign
func GetH5PaySign(appId, nonceStr, packages, signType, timeStamp, apiKey string) (paySign string) {
	// 原始字符串
	raw := fmt.Sprintf("appId=%s&nonceStr=%s&package=%s&signType=%s&timeStamp=%s&key=%s",
		appId, nonceStr, packages, signType, timeStamp, apiKey)
	buffer := new(bytes.Buffer)
	buffer.WriteString(raw)
	signStr := buffer.String()
	// 加密签名
	paySign = SignWithType(signType, signStr, apiKey)
	return
}
