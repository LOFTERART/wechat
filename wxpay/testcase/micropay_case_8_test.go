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

// 刷卡支付-用例8：【刷卡-异常】订单金额0.34元（含0.01元代金券），微信支付返回超时
func TestMicropayCase8(t *testing.T) {
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
	body.Body = "刷卡支付-测试用例8"
	body.OutTradeNo = outTradeNo
	body.TotalFee = 34
	body.SpbillCreateIP = "1.1.1.1"
	body.AuthCode = "150000111122223333"
	// 请求支付
	wxRsp, err = testClient.Micropay(body)
	if err != nil {
		return
	}
	// FIXME 校验字段
	// if err = CheckFields(wxRsp, map[string]interface{}{
	// 	"result_code": wxpay.ResponseFail,
	// 	"err_code":    wxpay.ErrCodeSystemError,
	// }); err != nil {
	// 	return
	// }
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
		"result_code": wxpay.ResponseSuccess,
		"recall":      "N",
	}); err != nil {
		return
	}
	// 更换商户订单号，重新发起支付
	outTradeNo = util.RandomString(32)
	body.OutTradeNo = outTradeNo
	wxRsp, err = testClient.Micropay(body)
	if err != nil {
		return
	}
	if err = CheckFields(wxRsp, map[string]interface{}{
		"return_code": wxpay.ResponseSuccess,
		"err_code":    wxpay.ResponseSuccess,
		"coupon_fee":  1,
	}); err != nil {
		return
	}
}
