package wechat

import (
	"fmt"
	"testing"
)

// 测试获取访问凭证
func TestAccessToken(t *testing.T) {
	fmt.Println("----------获取访问凭证----------")
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

// 测试获取用户基本信息(UnionID机制)
func TestWeChatUserInfo(t *testing.T) {
	fmt.Println("----------获取用户基本信息----------")
	// 请求接口
	token := "25_p1C7uUEPdgWHqCgX3hNcdBNZqdjEKU75ZLGQKfUFjEq7mQms-8J1KmSD0Fh0wSBg1pwumB_kRoB8OnnR10fml914bazh9xAoigZT2QHQrentZpZ--SM2j3iKaMRgr0Ec9_xjFKrPw_N5Og4mGDHeAHAZWA"
	openId := "gh_2d95fca4a95e"
	user, err := GetWeChatUserInfo(token, openId, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", user)
}
