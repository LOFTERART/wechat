package wxapplet

import (
	"encoding/json"
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
)

// 登录凭证校验
func CodeToSession(appId string, secret string, jsCode string) (wxRsp CodeToSessionResponse, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, secret, jsCode)
	body, err := util.HttpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &wxRsp)
	return
}

type CodeToSessionResponse struct {
	OpenId     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionId    string `json:"unionid"`     // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	ErrCode    int    `json:"errcode"`     // 错误码
	ErrMsg     string `json:"errmsg"`      // 错误信息
}
