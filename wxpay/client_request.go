package wxpay

import (
	"encoding/json"
	"gitee.com/xiaochengtech/wechat/constant"
	"gitee.com/xiaochengtech/wechat/util"
	"io/ioutil"
	"strings"
)

type buildHandler func(bodyObj interface{}) (body map[string]interface{}, err error)

// 构建Body
func (c *Client) buildBody(bodyObj interface{}) (body map[string]interface{}, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	bodyJson, _ := json.Marshal(bodyObj)
	body = make(map[string]interface{})
	_ = json.Unmarshal(bodyJson, &body)
	// 添加固定参数
	body["appid"] = c.Config.AppId
	body["mch_id"] = c.Config.MchId
	if c.IsFacilitator() {
		body["sub_appid"] = c.Config.SubAppId
		body["sub_mch_id"] = c.Config.SubMchId
	}
	nonceStr := util.RandomString(32)
	body["nonce_str"] = nonceStr
	// 生成签名
	signType, _ := body["sign_type"].(string)
	var sign string
	if c.IsProd {
		sign = c.localSign(body, signType, c.apiKey)
	} else {
		body["sign_type"] = constant.SignTypeMD5
		key, iErr := c.sandboxSign(nonceStr, constant.SignTypeMD5)
		if err = iErr; iErr != nil {
			return
		}
		sign = c.localSign(body, constant.SignTypeMD5, key)
	}
	body["sign"] = sign
	return
}

// FIXME 用于微信找零的构建Body
func (c *Client) buildBodyInMchMode(bodyObj interface{}) (body map[string]interface{}, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	bodyJson, _ := json.Marshal(bodyObj)
	body = make(map[string]interface{})
	_ = json.Unmarshal(bodyJson, &body)
	// 添加固定参数
	body["mch_appid"] = c.Config.AppId
	body["mchid"] = c.Config.MchId
	if c.IsFacilitator() {
		body["sub_appid"] = c.Config.SubAppId
		body["sub_mch_id"] = c.Config.SubMchId
	}
	nonceStr := util.RandomString(32)
	body["nonce_str"] = nonceStr
	// 生成签名
	signType, _ := body["sign_type"].(string)
	var sign string
	if c.IsProd {
		sign = c.localSign(body, signType, c.apiKey)
	} else {
		body["sign_type"] = constant.SignTypeMD5
		key, iErr := c.sandboxSign(nonceStr, constant.SignTypeMD5)
		if err = iErr; iErr != nil {
			return
		}
		sign = c.localSign(body, constant.SignTypeMD5, key)
	}
	body["sign"] = sign
	return
}

// 向微信发送请求
func (c *Client) doWeChat(relativeUrl string, bodyObj interface{}) (bytes []byte, err error) {
	// 转换参数
	body, err := c.buildBody(bodyObj)
	if err != nil {
		return
	}
	// 发起请求
	bytes, err = util.HttpPostXml(c.URL(relativeUrl), util.GenerateXml(body))
	return
}

// 向微信发送带证书请求
func (c *Client) doWeChatWithCert(relativeUrl string, bodyObj interface{}, handle buildHandler) (bytes []byte, err error) {
	// 转换参数
	if handle == nil {
		handle = c.buildBody
	}
	body, err := handle(bodyObj)
	if err != nil {
		return
	}
	// 设置证书和连接池
	if err = c.setCertData(c.certFilepath); err != nil {
		return
	}
	// 发起请求
	resp, err := c.certClient.Post(c.URL(relativeUrl), "application/xml", strings.NewReader(util.GenerateXml(body)))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(resp.Body)
	return
}
