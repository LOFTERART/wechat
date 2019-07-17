package wechat

import (
	"encoding/xml"
)

// 授权码查询openid
func (c *Client) OpenIdByAuthCode(body OpenIdByAuthCodeBody) (wxRsp OpenIdByAuthCodeResponse, err error) {
	// 业务逻辑
	bytes, err := c.doWeChat("tools/authcodetoopenid", body)
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

// 授权码查询openid参数
type OpenIdByAuthCodeBody struct {
	SignType string `json:"sign_type,omitempty"` // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	AuthCode string `json:"auth_code"`           // 扫码支付授权码，设备读取用户微信中的条码或者二维码信息
}

// 授权码查询openid返回值
type OpenIdByAuthCodeResponse struct {
	ResponseModel
	// 当return_code为SUCCESS时
	ServiceResponseModel
	// 当return_code 和result_code都为SUCCESS时
	OpenId    string `xml:"openid"`     // 用户在商户appid下的唯一标识
	SubOpenId string `xml:"sub_openid"` // 用户在子商户appid下的唯一标识
}
