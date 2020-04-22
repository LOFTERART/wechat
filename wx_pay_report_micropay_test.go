package wechat

import (
	"fmt"
	"testing"
	"time"
)

// TODO 测试交易保障(MICROPAY)
func testReportMicropay(t *testing.T) {
	fmt.Println("----------交易保障(MICROPAY)----------")
	// 初始化参数
	body := ReportMicropayBody{}
	body.UserIp = "8.8.8.8"
	body.Trades = []ReportMicropayBodyTrade{
		{
			OutTradeNo: "vk4264I1UQ3Hm3E4AKsavK8npylGSgQA092f9ckUxp8A2gXmnsLEdsupURVTcaC7",
			BeginTime:  FormatDateTime(time.Now()),
			EndTime:    FormatDateTime(time.Now()),
			State:      ReportMicropayTradeStateOk,
		},
	}
	// 请求交易保障
	wxRsp, err := testClient.ReportMicropay(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
