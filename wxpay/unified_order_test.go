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

func TestUnifiedOrder(t *testing.T) {
	testUnifiedOrder(t)
}

// 测试统一下单
func testUnifiedOrder(t *testing.T) (outTradeNo string) {
	fmt.Println("----------统一下单----------")
	// 初始化参数
	body := UnifiedOrderBody{}
	body.Body = "7克拉车场-京TTT001-微信支付-停车费"
	body.OutTradeNo = "wxcs201908271600001111"
	body.TotalFee = 301
	body.SpbillCreateIP = "124.77.173.62"
	body.NotifyUrl = "http://www.gopay.ink"
	body.TradeType = TradeTypeNative
	// 请求支付
	wxRsp, err := testClient.UnifiedOrder(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
	return
}
