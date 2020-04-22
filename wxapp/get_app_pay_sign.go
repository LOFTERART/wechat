package wxapp

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
	"sort"
)

// APP支付，统一下单获取支付参数后，再次计算APP支付所需要的的sign
func GetAppPaySign(appId, nonceStr, partnerId, prepayId, signType, timeStamp, apiKey string) (paySign string) {
	// 原始字符串
	raw := fmt.Sprintf("appId=%s&nonceStr=%s&package==Sign=WXPay&partnerid=%s&prepayid=%s&timeStamp=%s&key=%s",
		appId, nonceStr, partnerId, prepayId, timeStamp, apiKey)
	buffer := new(bytes.Buffer)
	buffer.WriteString(raw)
	// 加密签名
	signStr := buffer.String()
	paySign = util.SignWithType(signType, signStr, apiKey)
	return
}

// 生成JS-SDK权限验证的签名
func GetTicketSign(nonceStr, ticket, timeStamp, url string) (ticketSign string) {
	// 生成参数排序并拼接
	signStr := sortSignParams(nonceStr, ticket, timeStamp, url)
	// 加密签名
	ticketSign = hex.EncodeToString(util.Sha1(signStr))
	return
}

// 获取根据Key排序后的请求参数字符串
func sortSignParams(nonceStr, ticket, timeStamp, url string) string {
	body := make(map[string]interface{})
	body["noncestr"] = nonceStr
	body["jsapi_ticket"] = ticket
	body["timestamp"] = timeStamp
	body["url"] = url
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
	return buffer.String()
}
