package wechat

import (
	"fmt"
	"testing"
)

// 测试查询订单
func testQueryOrder(t *testing.T, outTradeNo string) {
	fmt.Println("----------查询订单----------")
	// 初始化参数
	body := QueryOrderBody{}
	body.OutTradeNo = outTradeNo
	// 请求订单查询
	wxRsp, err := testClient.QueryOrder(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}

func TestQueryOrder(t *testing.T) {
	testQueryOrder(t, "wxcs201908231600001111")
}
