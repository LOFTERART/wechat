package wechat

// 微信支付的整体配置
type Config struct {
	AppId    string // 微信分配的公众账号ID
	SubAppId string // 微信分配的子商户公众账号ID
	MchId    string // 微信支付分配的商户号
	SubMchId string // 微信支付分配的子商户号，开发者模式下必填
}
