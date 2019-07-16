package wechat

import "encoding/xml"

// 关闭订单
// 境内普通商户：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
// 境内的服务商：https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_3
func (c *Client) CloseOrder(body CloseOrderBody) (wxRsp CloseOrderResponse, err error) {
	bytes, err := c.doWeChat("pay/closeorder", body)
	if err != nil {
		return
	}
	err = xml.Unmarshal(bytes, &wxRsp)
	return
}

// 关闭订单的参数
type CloseOrderBody struct {
	SignType   string `json:"sign_type,omitempty"` // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	OutTradeNo string `json:"out_trade_no"`        // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。详见商户订单号
}

// 关闭订单的返回值
type CloseOrderResponse struct {
	ResponseModel
	// 当return_code为SUCCESS时
	ServiceResponseModel
	ResultMsg string `xml:"result_msg"` // 对业务结果的补充说明
}
