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

type NotifyRefundHandler func(NotifyRefundBody) (NotifyResponseModel, error)

// 退款结果通知
func (c *Client) NotifyRefund(handler NotifyRefundHandler, requestBody []byte) (rspBody string, err error) {
	// TODO
	return
}

// 退款结果通知的参数
type NotifyRefundBody struct {
	ResponseModel
	// 当return_code为SUCCESS时
	ServiceResponseModel
	ReqInfo string `xml:"req_info"` // 加密信息请用商户秘钥进行解密，详见解密方式
	// 返回的加密字段
	TransactionId       string `xml:"transaction_id"`        // 微信订单号
	OutTradeNo          string `xml:"out_trade_no"`          // 商户系统内部的订单号
	RefundId            string `xml:"refund_id"`             // 微信退款单号
	OutRefundNo         string `xml:"out_refund_no"`         // 商户退款单号
	TotalFee            int    `xml:"total_fee"`             // 订单总金额，单位为分，只能为整数，详见支付金额
	SettlementTotalFee  int    `xml:"settlement_total_fee"`  // 当该订单有使用非充值券时，返回此字段。应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	RefundFee           int    `xml:"refund_fee"`            // 退款总金额,单位为分
	SettlementRefundFee int    `xml:"settlement_refund_fee"` // 退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	RefundStatus        string `xml:"refund_status"`         // 退款状态(RefundStatusXXX)
	SuccessTime         string `xml:"success_time"`          // 资金退款至用户帐号的时间，格式2017-12-15 09:46:01
	RefundRecvAccount   string `xml:"refund_recv_accout"`    // TODO 取当前退款单的退款入账方 1）退回银行卡：{银行名称}{卡类型}{卡尾号} 2）退回支付用户零钱: 支付用户零钱 3）退还商户: 商户基本账户 商户结算银行账户 4）退回支付用户零钱通: 支付用户零钱通
	RefundAccount       string `xml:"refund_account"`        // 退款资金来源(RefundAccountXXX)
	RefundRequestSource string `xml:"refund_request_source"` // 退款发起来源(RefundRequestSourceXXX)
}
