package wxpay

import (
	"net/http"
)

// 微信支付的整体配置
type Config struct {
	AppId    string // 微信分配的公众账号ID
	SubAppId string // 微信分配的子商户公众账号ID
	MchId    string // 微信支付分配的商户号
	SubMchId string // 微信支付分配的子商户号，开发者模式下必填
}

// 微信支付客户端配置
type Client struct {
	// 可公开参数
	Config      Config // 配置信息
	ServiceType int    // 服务模式
	IsProd      bool   // 是否是生产环境
	// 保密参数
	apiKey       string       // API Key
	certFilepath string       // 证书目录
	certClient   *http.Client // 带证书的http连接池
}

// 初始化微信支付客户端
func NewClient(isProd bool, serviceType int, apiKey string, certFilepath string, config Config) (client *Client) {
	client = new(Client)
	client.Config = config
	client.ServiceType = serviceType
	client.IsProd = isProd
	client.apiKey = apiKey
	client.certFilepath = certFilepath
	return client
}

// 是否是服务商模式
func (c *Client) IsFacilitator() bool {
	switch c.ServiceType {
	case ServiceTypeFacilitatorDomestic, ServiceTypeFacilitatorAbroad, ServiceTypeBankServiceProvidor:
		return true
	default:
		return false
	}
}

// 拼接完整的URL
func (c *Client) URL(relativePath string) string {
	if c.IsProd {
		return BaseUrl + relativePath
	} else {
		return BaseUrlSandbox + relativePath
	}
}
