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

package util

import (
	"encoding/hex"
	"gitee.com/xiaochengtech/wechat/constant"
	"strings"
)

// 根据签名类型，生成签名
func SignWithType(signType string, origin string, apiKey string) string {
	var hashSign []byte
	if signType == constant.SignTypeHmacSHA256 {
		hashSign = HmacSha256(origin, apiKey)
	} else {
		hashSign = Md5(origin)
	}
	return strings.ToUpper(hex.EncodeToString(hashSign))
}
