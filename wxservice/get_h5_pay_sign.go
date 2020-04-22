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

package wxservice

import (
	"bytes"
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
)

// JSAPI支付，统一下单获取支付参数后，再次计算出微信内H5支付需要用的paySign
func GetH5PaySign(appId, nonceStr, packages, signType, timeStamp, apiKey string) (paySign string) {
	// 原始字符串
	raw := fmt.Sprintf("appId=%s&nonceStr=%s&package=%s&signType=%s&timeStamp=%s&key=%s",
		appId, nonceStr, packages, signType, timeStamp, apiKey)
	buffer := new(bytes.Buffer)
	buffer.WriteString(raw)
	signStr := buffer.String()
	// 加密签名
	paySign = util.SignWithType(signType, signStr, apiKey)
	return
}
