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
)

// 刷卡支付-用例1：【刷卡-正常】订单金额0.01元，用户支付成功
func TestMicropayCase1(t *testing.T) {
	var (
		wxRsp wxpay.MicropayResponse
		err   error
	)
	defer func() {
		if err != nil {
			t.Logf("返回值: %+v\n", wxRsp)
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
	body.SceneInfo = &wxpay.SceneInfoModel{
		ID:   "1",
		Name: "测试门店",
	}
	// 请求支付
	wxRsp, err = testClient.Micropay(body)
	if err != nil {
		return
	}
	// 校验字段(状态)
	if wxRsp.ReturnCode != wxpay.ResponseSuccess {
		err = errors.New("return_code错误")
		return
	}
	if wxRsp.ReturnMsg != wxpay.ResponseMessageOk {
		err = errors.New("return_msg错误")
		return
	}
	if wxRsp.ResultCode != wxpay.ResponseSuccess {
		err = errors.New("result_code错误")
		return
	}
	if wxRsp.ErrCode != wxpay.ResponseSuccess {
		err = errors.New("err_code错误")
		return
	}
	// 校验字段(业务)
	if wxRsp.TradeType != wxpay.TradeTypeMicropay {
		err = errors.New("trade_type错误")
		return
	}
	if wxRsp.CashFee != 1 {
		err = errors.New("cash_fee错误")
		return
	}
	if wxRsp.CashFeeType != wxpay.FeeTypeCNY {
		err = errors.New("cash_fee_type错误")
		return
	}
	if wxRsp.OutTradeNo != outTradeNo {
		err = errors.New("out_trade_no错误")
		return
	}
	if wxRsp.TotalFee != 1 {
		err = errors.New("total_fee错误")
		return
	}
	return
}
