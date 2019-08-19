package wechat

type Client struct {
	config       Config // 配置信息
	serviceType  int    // 服务模式
	apiKey       string // API Key
	certFilepath string // 证书目录
	certData     []byte // 证书内容
	isProd       bool   // 是否是生产环境
}

// 是否是服务商模式
func (c *Client) isFacilitator() bool {
	switch c.serviceType {
	case ServiceTypeFacilitatorDomestic, ServiceTypeFacilitatorAbroad, ServiceTypeBankServiceProvidor:
		return true
	default:
		return false
	}
}

// 拼接完整的URL
func (c *Client) url(relativePath string) string {
	if c.isProd {
		return baseUrl + relativePath
	} else {
		return baseUrlSandbox + relativePath
	}
}

// 初始化微信客户端
func NewClient(isProd bool, serviceType int, apiKey string, certFilepath string, config Config) (client *Client) {
	client = new(Client)
	client.config = config
	client.serviceType = serviceType
	client.apiKey = apiKey
	client.certFilepath = certFilepath
	client.isProd = isProd
	return client
}

// // 撤销订单
// // 文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
// func (c *WeChatClient) Reverse(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatReverseResponse, err error) {
// 	var bytes []byte
// 	if c.isProd {
// 		pkcsPool := x509.NewCertPool()
// 		pkcs, iErr := ioutil.ReadFile(pkcs12FilePath)
// 		if err = iErr; iErr != nil {
// 			return
// 		}
// 		pkcsPool.AppendCertsFromPEM(pkcs)
// 		certificate, iErr := tls.LoadX509KeyPair(certFilePath, keyFilePath)
// 		if err = iErr; iErr != nil {
// 			return
// 		}
// 		tlsConfig := new(tls.Config)
// 		tlsConfig.Certificates = []tls.Certificate{certificate}
// 		tlsConfig.RootCAs = pkcsPool
// 		tlsConfig.InsecureSkipVerify = true
// 		if bytes, err = c.doWeChat(body, wxURL_Reverse, tlsConfig); err != nil {
// 			return
// 		}
// 	} else {
// 		if bytes, err = c.doWeChat(body, wxURL_SanBox_Reverse); err != nil {
// 			return
// 		}
// 	}
// 	err = xml.Unmarshal(bytes, &wxRsp)
// 	return
// }
//
// // 下载资金账单
// // 文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7
// // 好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
// func (c *WeChatClient) DownloadFundFlow(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
// 	var bytes []byte
// 	if c.isProd {
// 		pkcsPool := x509.NewCertPool()
// 		pkcs, iErr := ioutil.ReadFile(pkcs12FilePath)
// 		if err = iErr; iErr != nil {
// 			return
// 		}
// 		pkcsPool.AppendCertsFromPEM(pkcs)
// 		certificate, iErr := tls.LoadX509KeyPair(certFilePath, keyFilePath)
// 		if err = iErr; iErr != nil {
// 			return
// 		}
// 		tlsConfig := new(tls.Config)
// 		tlsConfig.Certificates = []tls.Certificate{certificate}
// 		tlsConfig.RootCAs = pkcsPool
// 		tlsConfig.InsecureSkipVerify = true
// 		bytes, err = c.doWeChat(body, wxURL_DownloadFundFlow, tlsConfig)
// 	} else {
// 		bytes, err = c.doWeChat(body, wxURL_SanBox_DownloadFundFlow)
// 	}
// 	wxRsp = string(bytes)
// 	return
// }
//
// // 拉取订单评价数据
// // 文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=11
// // 好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
// func (c *WeChatClient) BatchQueryComment(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
// 	var bytes []byte
// 	if c.isProd {
// 		body.Set("sign_type", SignTypeHmacSHA256)
// 		pkcsPool := x509.NewCertPool()
// 		pkcs, iErr := ioutil.ReadFile(pkcs12FilePath)
// 		if err = iErr; iErr != nil {
// 			return
// 		}
// 		pkcsPool.AppendCertsFromPEM(pkcs)
// 		certificate, iErr := tls.LoadX509KeyPair(certFilePath, keyFilePath)
// 		if err = iErr; iErr != nil {
// 			return
// 		}
// 		tlsConfig := new(tls.Config)
// 		tlsConfig.Certificates = []tls.Certificate{certificate}
// 		tlsConfig.RootCAs = pkcsPool
// 		tlsConfig.InsecureSkipVerify = true
// 		bytes, err = c.doWeChat(body, wxURL_BatchQueryComment, tlsConfig)
// 	} else {
// 		bytes, err = c.doWeChat(body, wxURL_SanBox_BatchQueryComment)
// 	}
// 	wxRsp = string(bytes)
// 	return
// }
