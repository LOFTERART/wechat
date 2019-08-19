package wechat

import (
	"fmt"
	"testing"
)

// 测试授权码查询openid
func testOpenIdByAuthCode(t *testing.T) {
	fmt.Println("----------授权码查询openid----------")
	// 初始化参数
	body := OpenIdByAuthCodeBody{}
	body.AuthCode = "135127679952609396"
	// 请求支付
	wxRsp, err := testClient.OpenIdByAuthCode(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
