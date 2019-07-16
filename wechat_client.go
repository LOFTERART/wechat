package wechat

type Client struct {
	config      Config // 配置信息
	serviceType int    // 服务模式
	apiKey      string // API Key
	isProd      bool   // 是否是生产环境
}

// 是否是服务商模式
func (c *Client) isFacilitator() bool {
	switch c.serviceType {
	case ServiceTypeFacilitator:
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
func NewClient(isProd bool, serviceType int, apiKey string, config Config) (client *Client) {
	client = new(Client)
	client.config = config
	client.serviceType = serviceType
	client.apiKey = apiKey
	client.isProd = isProd
	return client
}

//
//
// // 统一下单
// // 境内普通商户：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
// // 境内的服务商：https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_1
// func (c *WeChatClient) UnifiedOrder(body BodyMap) (wxRsp WeChatUnifiedOrderResponse, err error) {
// 	var bytes []byte
// 	if c.isProd {
// 		tlsConfig := new(tls.Config)
// 		tlsConfig.InsecureSkipVerify = true
// 		if bytes, err = c.doWeChat(body, wxUrlUnifiedOrder, tlsConfig); err != nil {
// 			return
// 		}
// 	} else {
// 		if bytes, err = c.doWeChat(body, wxUrlUnifiedOrderSandBox); err != nil {
// 			return
// 		}
// 	}
// 	err = xml.Unmarshal(bytes, &wxRsp)
// 	return
// }
//
// // 查询订单
// // 境内普通商户：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
// // 境内的服务商：https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_2
// func (c *WeChatClient) QueryOrder(body BodyMap) (wxRsp WeChatQueryOrderResponse, err error) {
// 	var bytes []byte
// 	if c.isProd {
// 		tlsConfig := new(tls.Config)
// 		tlsConfig.InsecureSkipVerify = true
// 		if bytes, err = c.doWeChat(body, wxUrlOrderQuery, tlsConfig); err != nil {
// 			return
// 		}
// 	} else {
// 		if bytes, err = c.doWeChat(body, wxUrlOrderQuerySandBox); err != nil {
// 			return
// 		}
// 	}
// 	err = xml.Unmarshal(bytes, &wxRsp)
// 	return
// }
//
// // 关闭订单
// // 文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
// func (c *WeChatClient) CloseOrder(body BodyMap) (wxRsp *WeChatCloseOrderResponse, err error) {
// 	var bytes []byte
// 	if c.isProd {
// 		tlsConfig := new(tls.Config)
// 		tlsConfig.InsecureSkipVerify = true
// 		if bytes, err = c.doWeChat(body, wxURL_CloseOrder, tlsConfig); err != nil {
// 			return
// 		}
// 	} else {
// 		if bytes, err = c.doWeChat(body, wxURL_SanBox_CloseOrder); err != nil {
// 			return
// 		}
// 	}
// 	err = xml.Unmarshal(bytes, &wxRsp)
// 	return
// }
//
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
// // 申请退款
// // 文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
// func (c *WeChatClient) Refund(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatRefundResponse, err error) {
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
// 		if bytes, err = c.doWeChat(body, wxURL_Refund, tlsConfig); err != nil {
// 			return
// 		}
// 	} else {
// 		if bytes, err = c.doWeChat(body, wxURL_SanBox_Refund); err != nil {
// 			return
// 		}
// 	}
// 	err = xml.Unmarshal(bytes, &wxRsp)
// 	return
// }
//
// // 查询退款
// // 文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
// func (c *WeChatClient) QueryRefund(body BodyMap) (wxRsp *WeChatQueryRefundResponse, err error) {
// 	var bytes []byte
// 	if c.isProd {
// 		tlsConfig := new(tls.Config)
// 		tlsConfig.InsecureSkipVerify = true
// 		if bytes, err = c.doWeChat(body, wxURL_RefundQuery, tlsConfig); err != nil {
// 			return
// 		}
// 	} else {
// 		if bytes, err = c.doWeChat(body, wxURL_SanBox_RefundQuery); err != nil {
// 			return
// 		}
// 	}
// 	err = xml.Unmarshal(bytes, &wxRsp)
// 	return
// }
//
// // 下载对账单
// // 文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6
// func (c *WeChatClient) DownloadBill(body BodyMap) (wxRsp string, err error) {
// 	var bytes []byte
// 	if c.isProd {
// 		tlsConfig := new(tls.Config)
// 		tlsConfig.InsecureSkipVerify = true
// 		bytes, err = c.doWeChat(body, wxURL_DownloadBill, tlsConfig)
// 	} else {
// 		bytes, err = c.doWeChat(body, wxURL_SanBox_DownloadBill)
// 	}
// 	wxRsp = string(bytes)
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
