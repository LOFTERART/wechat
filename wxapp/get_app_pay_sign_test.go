/*
   Copyright 2020 XiaochengTech

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package wxapp

import (
	"fmt"
	"gitee.com/xiaochengtech/wechat/constant"
	"gitee.com/xiaochengtech/wechat/util"
	"testing"
	"time"
)

func TestGetAppPaySign(t *testing.T) {
	fmt.Println("----------获取app的paysign----------")
	// 请求接口
	appId := "wx80adf00e00fecc80"
	nonceStr := util.RandomString(32)
	partnerId := util.RandomString(32)
	prepayId := util.RandomString(32)
	signType := constant.SignTypeMD5
	timeStamp := string(time.Now().Unix())
	apiKey := "TJ0Rg25wM2AfFltah6XXg5PxNZoyV9D4"
	sign := GetAppPaySign(appId, nonceStr, partnerId, prepayId, signType, timeStamp, apiKey)
	t.Logf("返回值: %+v\n", sign)
}
