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

package testcase

import (
	"errors"
	"gitee.com/xiaochengtech/wechat/util"
	"gitee.com/xiaochengtech/wechat/wxpay"
	"testing"
	"time"
)

// 刷卡支付-用例7：【刷卡-异常】订单金额0.33元，微信支付返回超时，且查单失败
func TestMicropayCase7(t *testing.T) {
	var (
		wxRsp    wxpay.MicropayResponse
		queryRsp wxpay.QueryOrderResponse
		err      error
	)
	defer func() {
		if err != nil {
			t.Logf("返回值: %+v\n", wxRsp)
			t.Logf("查询值: %+v\n", queryRsp)
			t.Error(err)
		}
	}()
	// 初始化参数
	outTradeNo := util.RandomString(32)
	body := wxpay.MicropayBody{}
	body.Body = "刷卡支付-测试用例7"
	body.OutTradeNo = outTradeNo
	body.TotalFee = 33
	body.SpbillCreateIP = "1.1.1.1"
	body.AuthCode = "150000111122223333"
	// 请求支付
	wxRsp, err = testClient.Micropay(body)
	if err == nil {
		err = errors.New("应该是超时而不是返回结果")
		return
	}
	// 查询订单
	for i := 0; i < 3; i++ {
		queryRsp, err = testClient.QueryOrder(wxpay.QueryOrderBody{
			OutTradeNo: outTradeNo,
		})
		if err == nil {
			err = errors.New("查询订单应该是超时而不是返回结果")
			return
		}
		// 10秒轮询1次，3次则超时
		time.Sleep(10)
	}
	// 撤销订单
	reverseRsp, err := testClient.Reverse(wxpay.ReverseBody{
		OutTradeNo: outTradeNo,
	})
	if err != nil {
		return
	}
	if err = CheckFields(reverseRsp, map[string]interface{}{
		"recall":      "N",
		"result_code": wxpay.ResponseSuccess,
	}); err != nil {
		return
	}
}
