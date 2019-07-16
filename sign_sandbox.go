package wechat

import (
	"encoding/xml"
	"errors"
)

// 获取沙盒签名Key的返回值
type getSignKeyResponse struct {
	ResponseModel
	Retcode        string `xml:"retcode"`
	MchId          string `xml:"mch_id"`
	SandboxSignkey string `xml:"sandbox_signkey"`
}

// 获取沙盒的签名
func sandboxSign(mchId, nonceStr, apiKey, signType string) (key string, err error) {
	body := make(BodyMap)
	body["mch_id"] = mchId
	body["nonce_str"] = nonceStr
	// 计算沙箱参数Sign
	sanboxSign := localSign(body, signType, apiKey)
	// 沙箱环境：获取key后，重新计算Sign
	key, err = getSandBoxSignKey(mchId, nonceStr, sanboxSign)
	return
}

// 调用微信提供的接口获取SandboxSignkey
func getSandBoxSignKey(mchId, nonceStr, sign string) (key string, err error) {
	params := make(map[string]interface{})
	params["mch_id"] = mchId
	params["nonce_str"] = nonceStr
	params["sign"] = sign
	paramXml := generateXml(params)
	bytes, err := httpPost(baseUrlSandbox+"pay/getsignkey", paramXml)
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
