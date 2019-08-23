package wechat

import (
	"fmt"
	"testing"
)

// 测试付款码支付
func testMicropay(t *testing.T) (outTradeNo string, transactionId string) {
	fmt.Println("----------付款码支付----------")
	// 初始化参数
	body := MicropayBody{}
	body.Body = "测试车场微信付款码支付-停车费"
	body.OutTradeNo = "wxcs201908231600001111"
	body.TotalFee = 1
	body.SpbillCreateIP = "124.77.173.62"
	body.AuthCode = "134514543656378466"
	body.SceneInfo = &SceneInfoModel{
		ID:   "1",
		Name: "测试门店",
	}
	// 请求支付
	wxRsp, err := testClient.Micropay(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
	transactionId = wxRsp.TransactionId
	return
}

func TestMicropay(t *testing.T) {
	testMicropay(t)
}
