package wechat

import (
	"fmt"
	"testing"
)

// 测试查询订单
func TestQueryOrder(t *testing.T) {
	fmt.Println("----------查询订单----------")
	// 初始化参数
	body := QueryOrderBody{}
	body.OutTradeNo = testOutOrderNos[0]
	// 请求订单查询
	wxRsp, err := testClient.QueryOrder(body)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
