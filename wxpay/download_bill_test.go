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

// 测试下载对账单
func testDownloadBill(t *testing.T) {
	fmt.Println("----------下载对账单----------")
	// 初始化参数
	body := DownloadBillBody{}
	body.BillDate = "20190701"
	body.BillType = BillTypeAll
	// 请求下载对账单
	wxRsp, _, err := testClient.DownloadBill(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
