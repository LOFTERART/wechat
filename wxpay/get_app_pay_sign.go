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
	"bytes"
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
)

// 调起支付接口
func GetAppPaySign(
	appId string,
	nonceStr string,
	partnerId string,
	prepayId string,
	signType string,
	timeStamp string,
	apiKey string,
) (paySign string) {
	// 原始字符串
	raw := fmt.Sprintf("appid=%s&noncestr=%s&package=Sign=WXPay&partnerid=%s&prepayid=%s&timestamp=%s&key=%s",
		appId, nonceStr, partnerId, prepayId, timeStamp, apiKey)
	buffer := new(bytes.Buffer)
	buffer.WriteString(raw)
	// 加密签名
	signStr := buffer.String()
	paySign = util.SignWithType(signType, signStr, apiKey)
	return
}
