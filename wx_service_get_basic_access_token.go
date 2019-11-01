package wechat

import (
	"encoding/json"
	"fmt"
)

// 获取基础支持的AccessToken
func GetBasicAccessToken(appId, appSecret string) (accessToken AccessToken, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	body, err := httpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &accessToken)
	return
}
