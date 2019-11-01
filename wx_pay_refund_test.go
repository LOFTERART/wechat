package wechat

import (
	"fmt"
	"testing"
)

// 测试申请退款
func testRefund(t *testing.T, outTradeNo string, transactionId string) {
	fmt.Println("----------申请退款----------")
	outRefundNo := GetRandomString(32)
	// 初始化参数
	body := RefundBody{}
	if transactionId != "" {
		body.TransactionId = transactionId
	}
	body.OutTradeNo = outTradeNo
	body.OutRefundNo = outRefundNo
	body.TotalFee = 100 // 必须用100
	body.RefundFee = 1
	// 请求申请退款
	wxRsp, err := testClient.Refund(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
