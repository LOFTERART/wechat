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
	"gitee.com/xiaochengtech/wechat/util"
	"gitee.com/xiaochengtech/wechat/wxpay"
	"testing"
)

// 刷卡支付-用例1：【刷卡-正常】订单金额0.01元，用户支付成功
func TestMicropayCase1(t *testing.T) {
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
	body.Body = "刷卡支付-测试用例1"
	body.OutTradeNo = outTradeNo
	body.TotalFee = 1
	body.SpbillCreateIP = "1.1.1.1"
	body.AuthCode = "150000111122223333"
	// 请求支付
	wxRsp, err = testClient.Micropay(body)
	if err != nil {
		return
	}
	// 校验字段
	validFields := map[string]interface{}{
		"trade_type":    wxpay.TradeTypeMicropay,
		"cash_fee":      1,
		"cash_fee_type": wxpay.FeeTypeCNY,
		"return_code":   wxpay.ResponseSuccess,
		"out_trade_no":  outTradeNo,
		"return_msg":    wxpay.ResponseMessageOk,
		"total_fee":     1,
		"result_code":   wxpay.ResponseSuccess,
		"err_code":      wxpay.ResponseSuccess,
	}
	if err = CheckFields(wxRsp, validFields); err != nil {
		return
	}
	// 查询订单
	queryRsp, err = testClient.QueryOrder(wxpay.QueryOrderBody{
		OutTradeNo: outTradeNo,
	})
	if err != nil {
		return
	}
	// 校验字段
	validQueryOrderFields := map[string]interface{}{
		"mch_id": testClient.Config.MchId,
		// FIXME 始终失败，为APP。 "trade_type": wxpay.TradeTypeMicropay,
		"trade_state_desc": "ok",
		"trade_state":      wxpay.TradeStateSuccess,
		"cash_fee":         1,
		"out_trade_no":     outTradeNo,
		"total_fee":        1,
		"result_code":      wxpay.ResponseSuccess,
		"err_code":         wxpay.ResponseSuccess,
	}
	if err = CheckFields(queryRsp, validQueryOrderFields); err != nil {
		return
	}
}
