package wechat

import (
	"encoding/json"
	"fmt"
)

// 向微信发送请求
func (c *Client) doWeChat(bodyObj interface{}, url string) (bytes []byte, err error) {
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
	body["nonce_str"] = GetRandomString(32)
	// 生成参数
	if !c.isProd {
		body["sign_type"] = SignTypeMD5
		// 从微信接口获取SandBoxSignKey
		key, iErr := sandboxSign(c.config.MchId, body["nonce_str"].(string), c.apiKey, body["sign_type"].(string))
		if err = iErr; iErr != nil {
			return
		}
		sign = localSign(body, body["sign_type"].(string), key)
	} else {
		// 本地计算Sign
		sign = localSign(body, body["sign_type"].(string), c.apiKey)
	}
	body["sign"] = sign
	reqXML := generateXml(body)
	// 发起请求
	bytes, err = HttpPost(url, reqXML)
	return
}
