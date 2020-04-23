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

// 刷卡支付-用例4：【刷卡-异常】订单金额0.3元，用户未支付完成，微信支付返回FAIL
func TestMicropayCase4(t *testing.T) {
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
	body.Body = "刷卡支付-测试用例4"
	body.OutTradeNo = outTradeNo
	body.TotalFee = 30
	body.SpbillCreateIP = "1.1.1.1"
	body.AuthCode = "150000111122223333"
	// 请求支付
	wxRsp, err = testClient.Micropay(body)
	if err != nil {
		return
	}
	// 刷卡支付api返回FAIL，err_code为USERPAYING(用户支付中)
	validFields := map[string]interface{}{
		"err_code_des": "需要输入密码",
		"err_code":     wxpay.ErrCodeUserPaying,
	}
	if err = CheckFields(wxRsp, validFields); err != nil {
		return
	}
	// 查询订单
	for i := 0; i < 3; i++ {
		queryRsp, err = testClient.QueryOrder(wxpay.QueryOrderBody{
			OutTradeNo: outTradeNo,
		})
		if err != nil {
			return
		}
		// 查单api返回支付成功
		if err = CheckFields(queryRsp, map[string]interface{}{
			"result_code": wxpay.ResponseSuccess,
			"total_fee":   30,
			"trade_state": wxpay.TradeStateSuccess,
		}); err == nil {
			return
		}
		// 查单api明确返回订单支付失败或其他状态
		if err = CheckFields(queryRsp, map[string]interface{}{
			"result_code": wxpay.ResponseSuccess,
			"total_fee":   30,
			"trade_state": wxpay.TradeStateUserPaying,
		}); err == nil {
			return
		}
		// 10秒轮询1次，3次则超时
		time.Sleep(10)
	}
	err = errors.New("查询订单超时")
}
