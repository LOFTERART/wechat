package wechat

// 返回结果的通信标识
type ResponseBase struct {
	ErrCode int    `json:"errcode"` // 错误码
	ErrMsg  string `json:"errmsg"`  // 错误信息
}

type JsapiTicket struct {
	Ticket    string `json:"ticket"`     // 获取到的凭证
	ExpiresIn int64  `json:"expires_in"` // SessionKey超时时间（秒）
	ErrCode   int    `json:"errcode"`    // 错误码
	ErrMsg    string `json:"errmsg"`     // 错误信息
}

type SessionKey struct {
	OpenId     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionId    string `json:"unionid"`     // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	ErrCode    int    `json:"errcode"`     // 错误码
	ErrMsg     string `json:"errmsg"`      // 错误信息
}
