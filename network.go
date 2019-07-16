package wechat

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// 发送Post请求，参数是XML格式的字符串
func HttpPost(url string, xmlBody string) (body []byte, err error) {
	resp, err := http.Post(url, "application/xml", strings.NewReader(xmlBody))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
