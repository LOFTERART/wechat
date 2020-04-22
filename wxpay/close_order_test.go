package wxpay

import (
	"fmt"
	"testing"
)

// 测试关闭订单
func testCloseOrder(t *testing.T, outTradeNo string) {
	fmt.Println("----------关闭订单----------")
	// 初始化参数
	body := CloseOrderBody{}
	body.OutTradeNo = outTradeNo
	// 请求关闭订单
	wxRsp, err := testClient.CloseOrder(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
