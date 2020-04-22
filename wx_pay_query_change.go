package wechat

import (
	"encoding/xml"
)

// 企业付款到零钱的查询
func (c *PayClient) QueryChange(body QueryChangeBody) (wxRsp QueryChangeResponse, err error) {
	// 业务逻辑
	bytes, err := c.doWeChatWithCert("mmpaymkttransfers/gettransferinfo", body)
	if err != nil {
		return
	}
	// 不返回sign不需要校验
	// 解析返回值
	err = xml.Unmarshal(bytes, &wxRsp)
	return
}

// 微信找零查询的参数
type QueryChangeBody struct {
	PartnerTradeNo string `json:"partner_trade_no"` // 商户系统内部订单号
}

// 微信找零查询的返回值
type QueryChangeResponse struct {
	ResponseModel
	MchServiceResponseModel
	PartnerTradeNo string `xml:"partner_trade_no"` // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。详见商户订单号
	DetailId       string `xml:"detail_id"`        // 调用企业付款API时，微信系统内部产生的单号
	Status         string `xml:"status"`           // 转账状态
	Reason         string `xml:"reason"`           // 失败原因
	OpenId         string `xml:"openid"`           // 转账的openid
	TransferName   string `xml:"transfer_name"`    // 收款用户姓名
	PaymentAmount  int64  `xml:"payment_amount"`   // 付款金额单位为“分”
	TransferTime   string `xml:"transfer_time"`    // 发起转账的时间
	PaymentTime    string `xml:"payment_time"`     // 企业付款成功时间
	Desc           string `xml:"desc"`             // 企业付款备注
}
