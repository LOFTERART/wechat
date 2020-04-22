package util

import (
	"regexp"
)

// 是否是微信付款码，18位纯数字，以10、11、12、13、14、15开头
func IsValidAuthCode(authcode string) (ok bool) {
	pattern := "^1[0-5][0-9]{16}$"
	ok, _ = regexp.MatchString(pattern, authcode)
	return
}
