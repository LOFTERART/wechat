package wechat

import (
	"fmt"
	"testing"
)

// 测试付款码支付
func TestWeChatClientMicropay(t *testing.T) {
	client := NewWeChatClient(AppID, MchID, ApiKey, IsProd())
	outTradeNo := GetRandomString(32)
	// 初始化参数
	body := MicropayBody{}
	if IsFacilitator() {
		body.SubMchId = SubMchID
	}
	body.NonceStr = GetRandomString(32)
	body.Body = "扫用户付款码支付"
	body.OutTradeNo = outTradeNo
	body.TotalFee = 1
	body.SpbillCreateIP = "124.77.173.62"
	body.AuthCode = "123456789012345678"
	body.SignType = SignTypeMD5
	body.SceneInfo = SceneInfoObjModel{
		ID: "1",
	}
	// 请求支付
	wxRsp, err := client.Micropay(body)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("付款码支付: %+v\n", wxRsp)
	// 查询结果
	// queryBody := make(BodyMap)
	// if IsFacilitator() {
	// 	queryBody["sub_mch_id"] = SubMchID
	// }
	// queryBody["out_trade_no"] = outTradeNo
	// queryBody["nonce_str"] = GetRandomString(32)
	// queryBody["sign_type"] = SignTypeMD5
	// // 请求订单查询
	// wxQueryRsp, err := client.QueryOrder(queryBody)
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Printf("付款码支付查询: %+v\n", wxQueryRsp)
}
