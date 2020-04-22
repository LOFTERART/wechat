package wxapplet

import (
	"fmt"
	"testing"
)

func TestCodeToSession(t *testing.T) {
	fmt.Println("----------登录凭证校验----------")
	// 请求接口
	appId := ""
	secret := ""
	jsCode := ""
	rst, err := CodeToSession(appId, secret, jsCode)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", rst)
}
