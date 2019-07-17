package wechat

import (
	"errors"
	"github.com/beevik/etree"
)

// 验证微信返回的结果签名
func (c *Client) doVerifySign(xmlStr []byte) (err error) {
	// 生成XML文档
	doc := etree.NewDocument()
	if err = doc.ReadFromBytes(xmlStr); err != nil {
		return
	}
	// 遍历所有Tag，生成Map和Sign
	result, targetSign := make(map[string]interface{}), ""
	for _, elem := range doc.SelectElement("xml").ChildElements() {
		// 跳过空值
		if elem.Text() == "" || elem.Text() == "0" {
			continue
		}
		if elem.Tag != "sign" {
			result[elem.Tag] = elem.Text()
		} else {
			targetSign = elem.Text()
		}
	}
	// 获取签名类型
	signType := SignTypeMD5
	if result["sign_type"] != nil {
		signType = result["sign_type"].(string)
	}
	// 生成签名
	var sign string
	if c.isProd {
		sign = localSign(result, signType, c.apiKey)
	} else {
		key, iErr := sandboxSign(c.config.MchId, result["nonce_str"].(string), c.apiKey, SignTypeMD5)
		if err = iErr; iErr != nil {
			return
		}
		sign = localSign(result, SignTypeMD5, key)
	}
	// 验证
	if targetSign != sign {
		err = errors.New("签名无效")
	}
	return
}
