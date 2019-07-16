package wechat

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 测试统一下单
func TestUnifiedOrder(t *testing.T) {
	fmt.Println("----------统一下单----------")
	client := NewClient(false, ServiceType, ApiKey, Config{
		AppId:    AppID,
		SubAppId: SubAppID,
		MchId:    MchID,
		SubMchId: SubMchID,
	})
	outTradeNo := GetRandomString(32)
	// 初始化参数
	body := UnifiedOrderBody{}
	body.Body = "统一下单支付"
	body.OutTradeNo = outTradeNo
	body.TotalFee = 101
	body.SpbillCreateIP = "124.77.173.62"
	body.NotifyUrl = "http://www.gopay.ink"
	body.TradeType = TradeTypeJsApi
	body.DeviceInfo = "WEB"
	// 请求支付
	wxRsp, err := client.UnifiedOrder(body)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Response: %+v\n", wxRsp)
	// 获取小程序需要的支付签名
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	pac := "prepay_id=" + wxRsp.PrepayId
	paySign := GetMiniPaySign("wxbf1c916561ebb420", wxRsp.NonceStr, pac, SignTypeMD5, timeStamp, ApiKey)
	fmt.Printf("paySign: %s\n", paySign)
}
