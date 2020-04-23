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

package wxofficial

import (
	"fmt"
	"testing"
)

func TestRefreshWebAuthAccessToken(t *testing.T) {
	fmt.Println("----------刷新网页授权的access_token----------")
	// 请求接口
	appId := ""
	refresh := ""
	token, err := RefreshWebAuthAccessToken(appId, refresh)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", token)
}
