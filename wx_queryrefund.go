package wechat

import "encoding/xml"

// 查询退款
// 境内普通商户：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
// 境内的服务商：https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_5
func (c *Client) QueryRefund(body QueryRefundBody) (wxRsp QueryRefundResponse, err error) {
	bytes, err := c.doWeChat("pay/refundquery", body)
	if err != nil {
		return
	}
	err = xml.Unmarshal(bytes, &wxRsp)
	return
}

// 查询退款的参数
type QueryRefundBody struct {
	SignType      string `json:"sign_type,omitempty"`      // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	TransactionId string `json:"transaction_id,omitempty"` // (非必填，四选一) 微信订单号 查询的优先级是： refund_id > out_refund_no > transaction_id > out_trade_no
	OutTradeNo    string `json:"out_trade_no,omitempty"`   // (非必填，四选一) 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	OutRefundNo   string `json:"out_refund_no,omitempty"`  // (非必填，四选一) 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	RefundId      string `json:"refund_id,omitempty"`      // (非必填，四选一) 微信退款单号
	Offset        string `json:"offset,omitempty"`         // (非必填) 偏移量，当部分退款次数超过10次时可使用，表示返回的查询结果从这个偏移量开始取记录
}

// 查询退款的返回值
type QueryRefundResponse struct {
	ResponseModel
	// 当return_code为SUCCESS时
	ServiceResponseModel
	TransactionId        string `xml:"transaction_id"`          // 微信订单号
	OutTradeNo           string `xml:"out_trade_no"`            // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	TotalFee             int    `xml:"total_fee"`               // 订单总金额，单位为分，只能为整数，详见支付金额
	SettlementTotalFee   int    `xml:"settlement_total_fee"`    // 当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。
	FeeType              string `xml:"fee_type"`                // 订单金额货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFee              int    `xml:"cash_fee"`                // 现金支付金额，单位为分，只能为整数，详见支付金额
	RefundCount          int    `xml:"refund_count"`            // 当前返回退款笔数
	OutRefundNo0         string `xml:"out_refund_no_0"`         // TODO
	RefundId0            string `xml:"refund_id_0"`             // TODO
	RefundChannel0       string `xml:"refund_channel_0"`        // TODO
	TotalRefundCount     int    `xml:"total_refund_count"`      // 订单总共已发生的部分退款次数，当请求参数传入offset后有返回
	RefundFee0           int    `xml:"refund_fee_0"`            // TODO
	SettlementRefundFee0 int    `xml:"settlement_refund_fee_0"` // TODO
	CouponType00         string `xml:"coupon_type_0_0"`         // TODO
	CouponRefundFee0     int    `xml:"coupon_refund_fee_0"`     // TODO
	CouponRefundCount0   int    `xml:"coupon_refund_count_0"`   // TODO
	CouponRefundId00     string `xml:"coupon_refund_id_0_0"`    // TODO
	CouponRefundFee00    int    `xml:"coupon_refund_fee_0_0"`   // TODO
	RefundStatus0        string `xml:"refund_status_0"`         // TODO
	RefundAccount0       string `xml:"refund_account_0"`        // TODO
	RefundRecvAccout0    string `xml:"refund_recv_accout_0"`    // TODO
	RefundSuccessTime0   string `xml:"refund_success_time_0"`   // TODO
}
