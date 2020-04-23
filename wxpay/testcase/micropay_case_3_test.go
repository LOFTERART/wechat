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

// 刷卡支付-用例3：【刷卡-正常】订单金额0.03元（含0.01元代金券和0.02元免充值现金券），用户支付成功
// FIXME 查询订单时，返回值验签无效，去掉xml里面的cash_fee:0即可通过
func TestMicropayCase3(t *testing.T) {
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
	body.Body = "刷卡支付-测试用例3"
	body.OutTradeNo = outTradeNo
	body.TotalFee = 3
	body.SpbillCreateIP = "1.1.1.1"
	body.AuthCode = "150000111122223333"
	// 请求支付
	wxRsp, err = testClient.Micropay(body)
	if err != nil {
		return
	}
	// 校验字段
	validFields := map[string]interface{}{
		"coupon_fee":           3,
		"out_trade_no":         outTradeNo,
		"return_code":          wxpay.ResponseSuccess,
		"settlement_total_fee": 1,
		"cash_fee":             0,
		"total_fee":            3,
		"result_code":          wxpay.ResponseSuccess,
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
		"coupon_fee":           3,
		"settlement_total_fee": 1,
		"return_code":          wxpay.ResponseSuccess,
		"cash_fee":             0,
		"out_trade_no":         outTradeNo,
		"total_fee":            3,
		"coupon_count":         2,
		"result_code":          wxpay.ResponseSuccess,
	}
	if err = CheckFields(queryRsp, validQueryOrderFields); err != nil {
		return
	}
	if len(queryRsp.Coupons) != 2 {
		err = errors.New("优惠券数量错误")
		return
	}
	if err = CheckFields(queryRsp.Coupons[0], map[string]interface{}{
		"CouponFee":  1,
		"CouponType": wxpay.CouponTypeCash,
	}); err != nil {
		return
	}
	if err = CheckFields(queryRsp.Coupons[1], map[string]interface{}{
		"CouponFee":  2,
		"CouponType": wxpay.CouponTypeNoCash,
	}); err != nil {
		return
	}
}
