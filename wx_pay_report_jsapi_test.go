package wechat

import (
	"fmt"
	"testing"
)

// TODO 测试交易保障(JSAPI)
func testReportJsApi(t *testing.T) {
	fmt.Println("----------交易保障(JSAPI)----------")
	// 初始化参数
	body := ReportJsApiBody{}
	body.InterfaceUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	body.ExecuteTime = 101
	body.ReturnCode = ResponseSuccess
	body.ResultCode = ResponseSuccess
	body.UserIp = "8.8.8.8"
	// 请求交易保障
	wxRsp, err := testClient.ReportJsApi(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
