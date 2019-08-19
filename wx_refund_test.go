package wechat

import (
	"fmt"
	"testing"
)

// 测试申请退款
func TestRefund(t *testing.T) {
	fmt.Println("----------申请退款----------")
	outRefundNo := GetRandomString(32)
	// 初始化参数
	body := RefundBody{}
	body.TransactionId = "4208133673020190816164908595660"
	body.OutTradeNo = "YgENQFTovdeJdFouNyy3nFVOhGD6ZvPH"
	body.OutRefundNo = outRefundNo
	body.TotalFee = 100 //必须用100
	body.RefundFee = 1

	// 请求申请退款
	wxRsp, err := testClient.Refund(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
