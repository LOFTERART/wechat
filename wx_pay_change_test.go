package wechat

import (
	"fmt"
	"testing"
)

// 测试微信找零
func TestChange(t *testing.T) {
	fmt.Println("----------微信找零----------")
	// 初始化参数
	body := ChangeBody{}
	body.PartnerTradeNo = "1111111"
	body.OpenId = "111"
	body.CheckName = CheckNameTypeNoCheck
	body.Amount = 30
	body.Desc = "停车费找零"
	body.SpbillCreateIP = "124.77.173.62"
	// 请求的客户端必须是商户模式，且是特殊的商户接口
	changeClient := NewPayClient(true, true, ServiceTypeNormalDomestic, testApiKey, testCertPath, Config{
		AppId: testAppId, // 用子商户id设置
		MchId: testMchId, // 用子商户号设置
	})
	// 请求支付
	wxRsp, err := changeClient.Change(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", wxRsp)
	return
}
