package wechat

import "encoding/xml"

// 撤销订单
func (c *Client) Reverse(body ReverseBody) (wxRsp ReverseResponse, err error) {
	// 业务逻辑
	bytes, err := c.doWeChatWithCert("secapi/pay/reverse", body)
	if err != nil {
		return
	}
	// 结果校验
	if err = c.doVerifySign(bytes, true); err != nil {
		return
	}
	// 解析返回值
	err = xml.Unmarshal(bytes, &wxRsp)
	return
}

// 撤销订单的参数
type ReverseBody struct {
	TransactionId string `json:"transaction_id,omitempty"` // 微信支付订单号
	OutTradeNo    string `json:"out_trade_no"`             // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
}

// 撤销订单的返回值
type ReverseResponse struct {
	ResponseModel
	ServiceResponseModel
	Recall string `xml:"recall"` // 是否需要继续调用撤销，Y-需要，N-不需要
}
