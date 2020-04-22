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
	"encoding/json"
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
)

// 获取基础支持的AccessToken
func GetBasicAccessToken(appId, appSecret string) (accessToken AccessToken, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	body, err := util.HttpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &accessToken)
	return
}

type AccessToken struct {
	AccessToken  string `json:"access_token"`  // 获取到的凭证
	ExpiresIn    int64  `json:"expires_in"`    // SessionKey超时时间（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_tokenOpenId
	OpenId       string `json:"openid"`        // 用户唯一标识
	Scope        string `json:"scope"`         // 用户授权的作用域
	ErrCode      int    `json:"errcode"`       // 错误码
	ErrMsg       string `json:"errmsg"`        // 错误信息
}
