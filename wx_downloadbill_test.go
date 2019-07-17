package wechat

import (
	"fmt"
	"testing"
)

// 测试下载对账单
func TestDownloadBill(t *testing.T) {
	fmt.Println("----------下载对账单----------")
	// 初始化参数
	body := DownloadBillBody{}
	body.BillDate = "20190701"
	body.BillType = BillTypeAll
	// 请求下载对账单
	wxRsp, _, err := testClient.DownloadBill(body)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
