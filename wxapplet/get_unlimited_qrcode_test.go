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

package wxapplet

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetUnlimitedQrcode(t *testing.T) {
	fmt.Println("----------获取小程序二维码----------")
	// 请求接口
	body := GetUnlimitedQrcodeBody{
		AccessToken: "",
		Scene:       "TJ0Rg25wM2AfFltah6XXg5PxNZoyV9D4",
	}
	data, iErr, err := GetUnlimitedQrcode(body)
	if err != nil {
		t.Error(err)
		return
	}
	if iErr.ErrCode > 0 {
		t.Logf("错误: %+v\n", iErr)
		return
	}
	err = ioutil.WriteFile("/Users/shallot/1.jpg", data, 0666)
	t.Logf("返回文件: ~/1.jpg\n")
}
