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

package wxapp

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
	"sort"
)

// 生成JS-SDK权限验证的签名
func GetTicketSign(nonceStr, ticket, timeStamp, url string) (ticketSign string) {
	// 生成参数排序并拼接
	signStr := sortSignParams(nonceStr, ticket, timeStamp, url)
	// 加密签名
	ticketSign = hex.EncodeToString(util.Sha1(signStr))
	return
}

// 获取根据Key排序后的请求参数字符串
func sortSignParams(nonceStr, ticket, timeStamp, url string) string {
	body := make(map[string]interface{})
	body["noncestr"] = nonceStr
	body["jsapi_ticket"] = ticket
	body["timestamp"] = timeStamp
	body["url"] = url
	keyList := make([]string, 0)
	for k := range body {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		s := fmt.Sprintf("%s=%s&", k, fmt.Sprintf("%v", body[k]))
		buffer.WriteString(s)
	}
	return buffer.String()
}
