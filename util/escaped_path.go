package util

import (
	"net/url"
)

// 对URL进行Encode编码
func EscapedPath(u string) (path string, err error) {
	uriObj, err := url.Parse(u)
	if err != nil {
		return
	}
	path = uriObj.EscapedPath()
	return
}
