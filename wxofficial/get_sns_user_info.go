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
	"encoding/json"
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
)

// 网页授权-拉取用户信息(需scope为snsapi_userinfo)
func GetSnsUserInfo(accessToken string, openId string, lang string) (wxRsp UserInfoResponse, err error) {
	if len(lang) <= 0 {
		lang = "zh_CN"
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s", accessToken, openId, lang)
	body, err := util.HttpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &wxRsp)
	return
}
