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

func TestQueryOrder(t *testing.T) {
	testQueryOrder(t, "wxcs201908231600003333")
}

// 测试查询订单
func testQueryOrder(t *testing.T, outTradeNo string) {
	fmt.Println("----------查询订单----------")
	// 初始化参数
	body := QueryOrderBody{}
	body.OutTradeNo = outTradeNo
	// 请求订单查询
	wxRsp, err := testClient.QueryOrder(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
