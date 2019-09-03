package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/cuckoopark/common/basic"
)

// 获取全局唯一后台接口调用凭据
func GetAccessToken(appId, appSecret string) (accessToken *AccessToken, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	body, err := httpGet(url)
	if err != nil {
		return
	}
	// 反序列化json对象
	accessToken = new(AccessToken)
	err = json.Unmarshal(body, &accessToken)
	return
}

// 获取用户基本信息(UnionID机制)
func GetWeChatUserInfo(accessToken, openId string, lang ...string) (userInfo *WeChatUserInfo, err error) {
	language := basic.If(len(lang) > 0, lang[0], "zh_CN")
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=%s", accessToken, openId, language)
	body, err := httpGet(url)
	// 反序列化json对象
	userInfo = new(WeChatUserInfo)
	err = json.Unmarshal(body, &userInfo)
	return
}
