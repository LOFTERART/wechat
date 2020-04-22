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

import (
	"encoding/xml"
)

// 付款码查询OpenId
func (c *Client) AuthCodeToOpenId(body AuthCodeToOpenIdBody) (wxRsp AuthCodeToOpenIdResponse, err error) {
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

// 付款码查询OpenId参数
type AuthCodeToOpenIdBody struct {
	SignType string `json:"sign_type,omitempty"` // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	AuthCode string `json:"auth_code"`           // 扫码支付授权码，设备读取用户微信中的条码或者二维码信息
}

// 付款码查询OpenId返回值
type AuthCodeToOpenIdResponse struct {
	ResponseModel
	// 当return_code为SUCCESS时
	ServiceResponseModel
	// 当return_code 和result_code都为SUCCESS时
	OpenId    string `xml:"openid"`     // 用户在商户appid下的唯一标识
	SubOpenId string `xml:"sub_openid"` // 用户在子商户appid下的唯一标识
}
