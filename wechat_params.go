package wechat

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/parnurzeal/gorequest"
)

// 从微信提供的接口获取：SandboxSignKey
func getSandBoxSign(mchId, nonceStr, apiKey, signType string) (key string, err error) {
	body := make(BodyMap)
	body["mch_id"] = mchId
	body["nonce_str"] = nonceStr
	// 计算沙箱参数Sign
	sanboxSign := getLocalSign(body, signType, apiKey)
	// 沙箱环境：获取key后，重新计算Sign
	key, err = getSandBoxSignKey(mchId, nonceStr, sanboxSign)
	return
}

// 从微信提供的接口获取：SandboxSignkey
func getSandBoxSignKey(mchId, nonceStr, sign string) (key string, err error) {
	reqs := make(BodyMap)
	reqs["mch_id"] = mchId
	reqs["nonce_str"] = nonceStr
	reqs["sign"] = sign
	reqXml := generateXml(reqs)
	_, byteList, errorList := gorequest.New().
		Post(wxURL_SanBox_GetSignKey).
		Type("xml").
		SendString(reqXml).EndBytes()
	if len(errorList) > 0 {
		err = errorList[0]
		return
	}
	var keyResponse getSignKeyResponse
	if err = xml.Unmarshal(byteList, &keyResponse); err != nil {
		return
	}
	if keyResponse.ReturnCode == "FAIL" {
		err = errors.New(keyResponse.Retmsg)
		return
	}
	key = keyResponse.SandboxSignkey
	return
}

// 生成请求XML的Body体
func generateXml(bm BodyMap) (reqXml string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml>")
	for k, v := range bm {
		buffer.WriteString("<")
		buffer.WriteString(k)
		buffer.WriteString("><![CDATA[")
		valueStr := convert2String(v)
		buffer.WriteString(valueStr)
		buffer.WriteString("]]></")
		buffer.WriteString(k)
		buffer.WriteString(">")
	}
	buffer.WriteString("</xml>")
	reqXml = buffer.String()
	return
}
