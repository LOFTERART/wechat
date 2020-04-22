/*
   Copyright 2020 XiaochengTech

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package wxpay

import (
	"errors"
	"gitee.com/xiaochengtech/wechat/constant"
	"github.com/beevik/etree"
)

// 验证微信返回的结果签名
func (c *Client) doVerifySign(xmlStr []byte, breakWhenFail bool) (err error) {
	// 生成XML文档
	doc := etree.NewDocument()
	if err = doc.ReadFromBytes(xmlStr); err != nil {
		return
	}
	root := doc.SelectElement("xml")
	// 验证return_code
	retCode := root.SelectElement("return_code").Text()
	if retCode != ResponseSuccess && breakWhenFail {
		return
	}
	// 遍历所有Tag，生成Map和Sign
	result, targetSign := make(map[string]interface{}), ""
	for _, elem := range root.ChildElements() {
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
	signType := constant.SignTypeMD5
	if result["sign_type"] != nil {
		signType = result["sign_type"].(string)
	}
	// 生成签名
	var sign string
	if c.IsProd {
		sign = c.localSign(result, signType, c.apiKey)
	} else {
		key, iErr := c.sandboxSign(result["nonce_str"].(string), constant.SignTypeMD5)
		if err = iErr; iErr != nil {
			return
		}
		sign = c.localSign(result, constant.SignTypeMD5, key)
	}
	// 验证
	if targetSign != sign {
		err = errors.New("签名无效")
	}
	return
}
