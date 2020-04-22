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
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
)

// 小程序支付，统一下单获取支付参数后，再次计算出小程序用的paySign
func GetAppletPaySign(
	appId string,
	nonceStr string,
	prepayId string,
	signType string,
	timeStamp string,
	apiKey string,
) (paySign string) {
	// 原始字符串
	signStr := fmt.Sprintf("appId=%s&nonceStr=%s&package=%s&signType=%s&timeStamp=%s&key=%s",
		appId, nonceStr, prepayId, signType, timeStamp, apiKey)
	// 加密签名
	paySign = util.SignWithType(signType, signStr, apiKey)
	return
}
