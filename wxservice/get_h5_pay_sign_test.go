package wxservice

import (
	"fmt"
	"gitee.com/xiaochengtech/wechat/constant"
	"gitee.com/xiaochengtech/wechat/util"
	"testing"
	"time"
)

func TestGetH5PaySign(t *testing.T) {
	fmt.Println("----------获取H5的支付签名----------")
	// 请求接口
	appId := "wx80adf00e00fecc80"
	nonceStr := util.RandomString(32)
	packages := util.RandomString(32)
	signType := constant.SignTypeMD5
	timeStamp := string(time.Now().Unix())
	apiKey := "TJ0Rg25wM2AfFltah6XXg5PxNZoyV9D4"
	sign := GetH5PaySign(appId, nonceStr, packages, signType, timeStamp, apiKey)
	t.Logf("返回值: %+v\n", sign)
}
