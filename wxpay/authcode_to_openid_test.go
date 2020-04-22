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

// 测试付款码查询openid
func TestAuthCodeToOpenId(t *testing.T) {
	fmt.Println("----------付款码查询openid----------")
	// 初始化参数
	body := AuthCodeToOpenIdBody{}
	body.AuthCode = "134785902462927760"
	// 请求支付
	wxRsp, err := testClient.AuthCodeToOpenId(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
