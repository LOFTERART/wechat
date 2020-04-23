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

package wxofficial

import (
	"encoding/hex"
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
)

// JS-SDK使用权限-签名算法
func GetJsApiTicketSign(
	nonceStr string,
	ticket string,
	timeStamp string,
	url string,
) (ticketSign string) {
	signStr := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonceStr, timeStamp, url)
	ticketSign = hex.EncodeToString(util.Sha1(signStr))
	return
}
