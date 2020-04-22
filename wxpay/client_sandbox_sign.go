package wxpay

import (
	"encoding/xml"
	"errors"
	"gitee.com/xiaochengtech/wechat/util"
)

// 获取沙盒的签名
func (c *Client) sandboxSign(nonceStr string, signType string) (key string, err error) {
	body := make(map[string]interface{})
	body["mch_id"] = c.Config.MchId
	body["nonce_str"] = nonceStr
	// 计算沙箱参数Sign
	sanboxSign := c.localSign(body, signType, c.apiKey)
	// 沙箱环境：获取key后，重新计算Sign
	key, err = c.getSandBoxSignKey(nonceStr, sanboxSign)
	return
}

// (沙盒环境) 调用微信提供的接口获取验签密钥
func (c *Client) getSandBoxSignKey(nonceStr string, sign string) (key string, err error) {
	params := make(map[string]interface{})
	params["mch_id"] = c.Config.MchId
	params["nonce_str"] = nonceStr
	params["sign"] = sign
	paramXml := util.GenerateXml(params)
	bytes, err := util.HttpPostXml(c.URL("pay/getsignkey"), paramXml)
	if err != nil {
		return
	}
	var keyResponse getSignKeyResponse
	if err = xml.Unmarshal(bytes, &keyResponse); err != nil {
		return
	}
	if keyResponse.ReturnCode == ResponseFail {
		err = errors.New(keyResponse.RetMsg)
		return
	}
	key = keyResponse.SandboxSignkey
	return
}

// (沙盒环境) 获取验签密钥接口的返回值
type getSignKeyResponse struct {
	ResponseModel
	Retcode        string `xml:"retcode"`         // TODO 移除？
	MchId          string `xml:"mch_id"`          // 调用接口提交的商户号
	SandboxSignkey string `xml:"sandbox_signkey"` // 返回的沙盒密钥
}
