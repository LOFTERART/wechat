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

package wxpay

import (
	"fmt"
	"testing"
)

// 测试微信找零查询
func TestQueryChange(t *testing.T) {
	fmt.Println("----------微信找零查询----------")
	// 初始化参数
	body := QueryChangeBody{}
	body.PartnerTradeNo = "wxcs201908241600005555"
	// 请求的客户端
	changeClient := NewClient(true, ServiceTypeNormalDomestic, testApiKey, testCertPath, Config{
		AppId: testAppId, // 用子商户id设置
		MchId: testMchId, // 用子商户号设置
	})
	// 请求支付
	wxRsp, err := changeClient.QueryChange(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", wxRsp)
	return
}
