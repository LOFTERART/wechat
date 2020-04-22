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

// TODO 测试交易保障(JSAPI)
func testReportJsApi(t *testing.T) {
	fmt.Println("----------交易保障(JSAPI)----------")
	// 初始化参数
	body := ReportJsApiBody{}
	body.InterfaceUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	body.ExecuteTime = 101
	body.ReturnCode = ResponseSuccess
	body.ResultCode = ResponseSuccess
	body.UserIp = "8.8.8.8"
	// 请求交易保障
	wxRsp, err := testClient.ReportJsApi(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
