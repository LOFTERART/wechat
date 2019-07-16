package wechat

import (
	"fmt"
	"testing"
)

// 测试付款码支付
func TestMicropay(t *testing.T) {
	client := NewClient(IsProd, ServiceType, ApiKey, Config{
		AppId:    AppID,
		MchId:    MchID,
		SubMchId: SubMchID,
	})
	outTradeNo := GetRandomString(32)
	// 初始化参数
	body := MicropayBody{}
	body.NonceStr = GetRandomString(32)
	body.Body = "扫用户付款码支付"
	body.OutTradeNo = outTradeNo
	body.TotalFee = 1
	body.SpbillCreateIP = "124.77.173.62"
	body.AuthCode = "123456789012345678"
	body.SignType = SignTypeMD5
	body.SceneInfo = JsonString(SceneInfoModel{
		ID:   "1",
		Name: "测试门店",
	})
	// 请求支付
	wxRsp, err := client.Micropay(body)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("付款码支付: %+v\n", wxRsp)
}
