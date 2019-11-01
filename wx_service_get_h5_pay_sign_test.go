package wechat

import (
	"fmt"
	"testing"
	"time"
)

func TestGetH5PaySign(t *testing.T) {
	fmt.Println("----------获取H5的支付签名----------")
	// 请求接口
	appId := "wx80adf00e00fecc80"
	nonceStr := GetRandomString(32)
	packages := GetRandomString(32)
	signType := SignTypeMD5
	timeStamp := string(time.Now().Unix())
	apiKey := "TJ0Rg25wM2AfFltah6XXg5PxNZoyV9D4"
	sign := GetH5PaySign(appId, nonceStr, packages, signType, timeStamp, apiKey)
	t.Logf("返回值: %+v\n", sign)
}
