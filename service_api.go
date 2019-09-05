/*
 基础支持相关的接口。
*/

package wechat

import (
	"encoding/json"
	"fmt"
)

// 获取基础支持的access_token
func GetBasicAccessToken(appId, appSecret string) (accessToken AccessToken, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	body, err := httpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &accessToken)
	return
}

// 获取用户基本信息(UnionID机制)
func GetBasicUserInfo(accessToken, openId, lang string) (userInfo UserInfo, err error) {
	if len(lang) <= 0 {
		lang = "zh_CN"
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=%s", accessToken, openId, lang)
	body, err := httpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &userInfo)
	return
}
