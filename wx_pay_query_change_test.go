package wechat

import (
	"fmt"
	"testing"
)

// 测试微信找零查询
func TestQueryChange(t *testing.T) {
	fmt.Println("----------微信找零查询----------")
	// 初始化参数
	body := QueryChangeBody{}
	body.PartnerTradeNo = "wxcs201908241600005555"
	// 请求的客户端必须是商户模式，且是特殊的商户接口
	changeClient := NewPayClient(true, false, ServiceTypeNormalDomestic, testApiKey, testCertPath, Config{
		AppId: testAppId, // 用子商户id设置
		MchId: testMchId, // 用子商户号设置
	})
	// 请求支付
	wxRsp, err := changeClient.QueryChange(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", wxRsp)
	return
}
