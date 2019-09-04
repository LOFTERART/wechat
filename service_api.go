package wechat

import (
	"encoding/json"
	"fmt"
)

// 获取全局唯一后台接口调用凭据
func GetAccessToken(appId string, appSecret string) (accessToken AccessToken, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	body, err := httpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &accessToken)
	return
}

// 获取用户基本信息(UnionID机制)
func GetUserInfo(accessToken string, openId string, lang ...string) (userInfo UserInfo, err error) {
	var language string
	if len(lang) > 0 {
		language = lang[0]
	} else {
		language = "zh_CN"
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=%s", accessToken, openId, language)
	body, err := httpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &userInfo)
	return
}
