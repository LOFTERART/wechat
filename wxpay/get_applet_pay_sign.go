package wxpay

import (
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
)

// 小程序支付，统一下单获取支付参数后，再次计算出小程序用的paySign
func GetAppletPaySign(
	appId string,
	nonceStr string,
	prepayId string,
	signType string,
	timeStamp string,
	apiKey string,
) (paySign string) {
	// 原始字符串
	signStr := fmt.Sprintf("appId=%s&nonceStr=%s&package=%s&signType=%s&timeStamp=%s&key=%s",
		appId, nonceStr, prepayId, signType, timeStamp, apiKey)
	// 加密签名
	paySign = util.SignWithType(signType, signStr, apiKey)
	return
}
