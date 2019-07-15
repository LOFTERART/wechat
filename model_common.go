package wechat

// 公共参数信息
type CommonModel struct {
	AppId      string `json:"appid"`                  // 微信分配的公众账号ID(企业号corpid即为此appId)
	SubAppId   string `json:"sub_appid,omitempty"`    // (服务商模式，非必填) 微信分配的子商户公众账号ID
	MchId      string `json:"mch_id"`                 // 微信支付分配的商户号
	SubMchId   string `json:"sub_mch_id,omitempty"`   // (服务商模式) 微信支付分配的子商户号
	NonceStr   string `json:"nonce_str"`              // 随机字符串，不长于32位。推荐随机数生成算法
	Sign       string `json:"sign,omitempty"`         // 签名，详见签名生成算法
	SignType   string `json:"sign_type,omitempty"`    // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

// 签名方式
const (
	SignTypeMD5        = "MD5" // 默认
	SignTypeHmacSHA256 = "HMAC-SHA256"
)
