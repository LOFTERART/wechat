package wechat

import (
	"fmt"
	"testing"
)

// 测试关闭订单
func TestCloseOrder(t *testing.T) {
	fmt.Println("----------关闭订单----------")
	// 初始化参数
	body := CloseOrderBody{}
	body.OutTradeNo = "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ"
	// 请求关闭订单
	wxRsp, err := testClient.CloseOrder(body)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Response: %+v\n", wxRsp)
}
