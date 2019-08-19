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
	body.OutTradeNo = "YgENQFTovdeJdFouNyy3nFVOhGD6ZvPH"
	// 请求查询退款
	wxRsp, err := testClient.QueryRefund(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
