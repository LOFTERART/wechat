package wechat

import (
	"encoding/xml"
	"github.com/beevik/etree"
)

type PayNotifyHandler func(PayNotifyBody) error

// 支付结果通知
func (c *Client) OnPayNotify(handler PayNotifyHandler, requestBody []byte) (rspBody string, err error) {
	var body PayNotifyBody
	if err = xml.Unmarshal(requestBody, &body); err != nil {
		return
	}
	// TODO 验证Sign
	// 解析CouponCount的对应项
	if body.CouponCount > 0 {
		doc := etree.NewDocument()
		if err = doc.ReadFromBytes(requestBody); err != nil {
			return
		}
		for i := 0; i < body.CouponCount; i++ {
			m := NewCouponResponseModel(doc, "coupon_id_%d", "coupon_type_%d", "coupon_fee_%d", i)
			body.Coupons = append(body.Coupons, m)
		}
	}
	// 调用外部处理
	if err = handler(body); err != nil {
		return
	}
	// 返回处理结果
	rspModel := NotifyResponseModel{
		ReturnCode: ResponseSuccess,
		ReturnMsg:  ResponseMessageOk,
	}
	rspBody = rspModel.ToXmlString()
	return
}

// 支付结果通知的参数
type PayNotifyBody struct {
	ResponseModel
	// 当return_code为SUCCESS时
	ServiceResponseModel
	DeviceInfo         string `xml:"device_info" json:"device_info"`                   // 微信支付分配的终端设备号
	IsSubscribe        string `xml:"is_subscribe" json:"is_subscribe"`                 // 用户是否关注公众账号(机构商户不返回)
	SubIsSubscribe     string `xml:"sub_is_subscribe" json:"sub_is_subscribe"`         // (服务商模式) 用户是否关注子公众账号(机构商户不返回)
	Openid             string `xml:"openid" json:"openid"`                             // 用户在商户appid下的唯一标识
	SubOpenId          string `xml:"sub_openid" json:"sub_openid"`                     // (服务商模式) 用户在子商户appid下的唯一标识
	TradeType          string `xml:"trade_type" json:"trade_type"`                     // 交易类型
	BankType           string `xml:"bank_type" json:"bank_type"`                       // 银行类型，采用字符串类型的银行标识，银行类型见附表
	TotalFee           int    `xml:"total_fee" json:"total_fee"`                       // 订单总金额，单位为分
	FeeType            string `xml:"fee_type" json:"fee_type"`                         // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFee            int    `xml:"cash_fee" json:"cash_fee"`                         // 现金支付金额订单现金支付金额，详见支付金额
	CashFeeType        string `xml:"cash_fee_type" json:"cash_fee_type"`               // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	SettlementTotalFee int    `xml:"settlement_total_fee" json:"settlement_total_fee"` // 应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	CouponFee          int    `xml:"coupon_fee" json:"coupon_fee"`                     // 代金券或立减优惠金额<=订单总金额，订单总金额-代金券或立减优惠金额=现金支付金额，详见支付金额
	CouponCount        int    `xml:"coupon_count" json:"coupon_count"`                 // 代金券或立减优惠使用数量
	TransactionId      string `xml:"transaction_id" json:"transaction_id"`             // 微信支付订单号
	OutTradeNo         string `xml:"out_trade_no" json:"out_trade_no"`                 // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	Attach             string `xml:"attach" json:"attach"`                             // 商家数据包，原样返回
	TimeEnd            string `xml:"time_end" json:"time_end"`                         // 支付完成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	// 使用coupon_count的序号生成的优惠券项
	Coupons []CouponResponseModel `xml:"-" json:"-"`
}
