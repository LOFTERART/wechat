package wxpay

import (
	"fmt"
	"testing"
)

// 测试查询退款
func testQueryRefund(t *testing.T, outTradeNo string) {
	fmt.Println("----------查询退款----------")
	// 初始化参数结构体
	body := QueryRefundBody{}
	body.OutTradeNo = outTradeNo
	// 请求查询退款
	wxRsp, err := testClient.QueryRefund(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
