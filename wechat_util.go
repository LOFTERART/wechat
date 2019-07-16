package wechat

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// 向微信发送请求
func (c *Client) doWeChat(bodyObj interface{}, url string, tlsConfig ...*tls.Config) (bytes []byte, err error) {
	bodyJson, _ := json.Marshal(bodyObj)
	var body map[string]interface{}
	_ = json.Unmarshal(bodyJson, &body)
	fmt.Printf("%+v \n", body)
	var sign string
	body["appid"] = c.config.AppId
	body["mch_id"] = c.config.MchId
	if c.isFacilitator() {
		body["sub_appid"] = c.config.SubAppId
		body["sub_mch_id"] = c.config.SubMchId
	}
	// 生成参数
	if !c.isProd {
		body["sign_type"] = SignTypeMD5
		// 从微信接口获取SandBoxSignKey
		key, iErr := getSandBoxSign(c.config.MchId, body["nonce_str"].(string), c.apiKey, body["sign_type"].(string))
		if err = iErr; iErr != nil {
			return
		}
		sign = getLocalSign(body, body["sign_type"].(string), key)
	} else {
		// 本地计算Sign
		sign = getLocalSign(body, body["sign_type"].(string), c.apiKey)
	}
	body["sign"] = sign
	reqXML := generateXml(body)
	// 发起请求
	agent := gorequest.New()
	if c.isProd && tlsConfig != nil {
		agent.TLSClientConfig(tlsConfig[0])
	}
	agent.Post(url)
	agent.Type("xml")
	agent.SendString(reqXML)
	_, bytes, errs := agent.EndBytes()
	if len(errs) > 0 {
		err = errs[0]
	}
	return
}
