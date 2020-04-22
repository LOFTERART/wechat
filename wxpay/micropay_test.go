/*
   Copyright 2020 XiaochengTech

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package wxpay

import (
	"fmt"
	"testing"
)

func TestMicropay(t *testing.T) {
	testMicropay(t)
}

// 测试付款码支付
func testMicropay(t *testing.T) (outTradeNo string, transactionId string) {
	fmt.Println("----------付款码支付----------")
	//outTradeNo = GetRandomString(32)
	// 初始化参数
	body := MicropayBody{}
	body.Body = "7克拉车场-京TTT001-微信支付-停车费"
	//body.OutTradeNo = outTradeNo
	body.OutTradeNo = "wxcs201909051000001111"
	body.TotalFee = 1
	body.SpbillCreateIP = "124.77.173.62"
	body.AuthCode = "135007001630843683"
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
