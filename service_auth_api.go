/*
 网页授权相关的接口。
 https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
*/

package wechat

import (
	"encoding/json"
	"fmt"
)

// 获取网页授权的access_token
func GetAuthAccessToken(appId, appSecret, code string) (accessToken AccessToken, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", appId, appSecret, code)
	body, err := httpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &accessToken)
	return
}

// 刷新网页授权的access_token
func RefreshAuthAccessToken(appId, refreshToken string) (accessToken AccessToken, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s", appId, refreshToken)
	body, err := httpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &accessToken)
	return
}

// 获取用户基本信息(授权机制)
func GetAuthUserInfo(accessToken, openId, lang string) (userInfo UserInfo, err error) {
	if len(lang) <= 0 {
		lang = "zh_CN"
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s", accessToken, openId, lang)
	body, err := httpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &userInfo)
	return
}

// 检验网页授权的access_token是否有效
func CheckAuthAccessToken(accessToken, openId string) (resp ResponseBase, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s", accessToken, openId)
	body, err := httpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}
