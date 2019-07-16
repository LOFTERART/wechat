package wechat

import (
	"fmt"
	"testing"
)

// 测试查询退款
func TestQueryRefund(t *testing.T) {
	fmt.Println("----------查询退款----------")
	// 初始化参数结构体
	body := QueryRefundBody{}
	body.OutTradeNo = "vk4264I1UQ3Hm3E4AKsavK8npylGSgQA092f9ckUxp8A2gXmnsLEdsupURVTcaC7"
	// 请求查询退款
	wxRsp, err := testClient.QueryRefund(body)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Response: %+v\n", wxRsp)
}
