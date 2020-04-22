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
	"sort"
)

// 本地通过支付参数计算签名值
func (c *Client) localSign(body map[string]interface{}, signType string, apiKey string) string {
	signStr := c.sortSignParams(body, apiKey)
	return util.SignWithType(signType, signStr, apiKey)
}

// 获取根据Key排序后的请求参数字符串
func (c *Client) sortSignParams(body map[string]interface{}, apiKey string) string {
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
	buffer.WriteString(fmt.Sprintf("key=%s", apiKey))
	return buffer.String()
}
