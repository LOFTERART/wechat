package wechat

import (
	"fmt"
	"testing"
	"time"
)

func TestGetAppPaySign(t *testing.T) {
	fmt.Println("----------获取app的paysign----------")
	// 请求接口
	appId := "wx80adf00e00fecc80"
	nonceStr := GetRandomString(32)
	partnerId := GetRandomString(32)
	prepayId := GetRandomString(32)
	signType := SignTypeMD5
	timeStamp := string(time.Now().Unix())
	apiKey := "TJ0Rg25wM2AfFltah6XXg5PxNZoyV9D4"
	sign := GetAppPaySign(appId, nonceStr, partnerId, prepayId, signType, timeStamp, apiKey)
	t.Logf("返回值: %+v\n", sign)
}
