package wxservice

import (
	"fmt"
	"testing"
)

func TestGetBasicUserInfo(t *testing.T) {
	fmt.Println("----------获取用户基本信息----------")
	// 请求接口
	token := ""
	openId := ""
	user, err := GetBasicUserInfo(token, openId, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", user)
}
