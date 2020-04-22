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

package util

import (
	"regexp"
)

// 是否是微信付款码，18位纯数字，以10、11、12、13、14、15开头
func IsValidAuthCode(authcode string) (ok bool) {
	pattern := "^1[0-5][0-9]{16}$"
	ok, _ = regexp.MatchString(pattern, authcode)
	return
}
