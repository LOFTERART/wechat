package wxapplet

import (
	"fmt"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	fmt.Println("----------获取小程序全局唯一后台接口调用凭据----------")
	// 请求接口
	appId := "wx80adf00e00fecc80"
	appSecret := "fa1c98a5449e909129d08b10c1bbb415"
	token, err := GetAccessToken(appId, appSecret)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", token)
}
