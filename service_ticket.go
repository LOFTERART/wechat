/*
 基础支持相关的接口。
*/

package wechat

import (
	"encoding/json"
	"fmt"
)

// 获取jsapi_ticket
func GetJSAPITicket(access_token string) (ticket JsapiTicket, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", access_token)
	body, err := httpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &ticket)
	return
}
