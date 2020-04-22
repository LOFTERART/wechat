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

import "encoding/xml"

// 撤销订单
func (c *Client) Reverse(body ReverseBody) (wxRsp ReverseResponse, err error) {
	// 业务逻辑
	bytes, err := c.doWeChatWithCert("secapi/pay/reverse", body, nil)
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
