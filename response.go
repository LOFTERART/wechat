package wechat

type JsapiTicket struct {
	Ticket    string `json:"ticket"`     // 获取到的凭证
	ExpiresIn int64  `json:"expires_in"` // SessionKey超时时间（秒）
	ErrCode   int    `json:"errcode"`    // 错误码
	ErrMsg    string `json:"errmsg"`     // 错误信息
}
