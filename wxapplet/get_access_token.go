package wxapplet

import (
	"encoding/json"
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
)

// 获取小程序全局唯一后台接口调用凭据
func GetAccessToken(appId string, appSecret string) (wxRsp GetAccessTokenResponse, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	body, err := util.HttpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &wxRsp)
	return
}

// 获取小程序全局唯一后台接口调用凭据返回值
type GetAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`  // 获取到的凭证
	ExpiresIn    int64  `json:"expires_in"`    // SessionKey超时时间（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_tokenOpenId
	OpenId       string `json:"openid"`        // 用户唯一标识
	Scope        string `json:"scope"`         // 用户授权的作用域
	ErrCode      int    `json:"errcode"`       // 错误码
	ErrMsg       string `json:"errmsg"`        // 错误信息
}
