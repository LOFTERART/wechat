package wxpay

import (
	"fmt"
	"testing"
)

// 测试付款码查询openid
func TestAuthCodeToOpenId(t *testing.T) {
	fmt.Println("----------付款码查询openid----------")
	// 初始化参数
	body := AuthCodeToOpenIdBody{}
	body.AuthCode = "134785902462927760"
	// 请求支付
	wxRsp, err := testClient.AuthCodeToOpenId(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
