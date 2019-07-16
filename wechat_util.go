package wechat

import (
	"encoding/json"
	"encoding/xml"
)

// 向微信发送请求
func (c *Client) doWeChat(relativeUrl string, bodyObj interface{}, rsp interface{}) (err error) {
	url := c.url(relativeUrl)
	bodyJson, _ := json.Marshal(bodyObj)
	body := make(map[string]interface{})
	_ = json.Unmarshal(bodyJson, &body)
	body["appid"] = c.config.AppId
	body["mch_id"] = c.config.MchId
	if c.isFacilitator() {
		body["sub_appid"] = c.config.SubAppId
		body["sub_mch_id"] = c.config.SubMchId
	}
	body["nonce_str"] = GetRandomString(32)
	// 生成参数
	signType, _ := body["sign_type"].(string)
	var sign string
	if c.isProd {
		sign = localSign(body, signType, c.apiKey)
	} else {
		body["sign_type"] = SignTypeMD5
		key, iErr := sandboxSign(c.config.MchId, body["nonce_str"].(string), c.apiKey, SignTypeMD5)
		if err = iErr; iErr != nil {
			return
		}
		sign = localSign(body, SignTypeMD5, key)
	}
	body["sign"] = sign
	reqXML := generateXml(body)
	// 发起请求
	bytes, err := httpPost(url, reqXML)
	if err != nil {
		return
	}
	// 解析参数
	err = xml.Unmarshal(bytes, rsp)
	if err != nil {
		return
	}
	return
}
