package wechat

import (
	"fmt"
	"testing"
)

// 测试付款码支付
func TestMicropay(t *testing.T) {
	fmt.Println("----------付款码支付----------")
	outTradeNo := GetRandomString(32)
	// 初始化参数
	body := MicropayBody{}
	body.Body = "测试车场付款码支付-停车费"
	body.OutTradeNo = outTradeNo
	body.TotalFee = 1
	body.SpbillCreateIP = "124.77.173.62"
	body.AuthCode = "123456789012345678"
	body.SceneInfo = &SceneInfoModel{
		ID:   "1",
		Name: "测试门店",
	}
	// 请求支付
	wxRsp, err := testClient.Micropay(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
