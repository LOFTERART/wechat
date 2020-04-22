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

import "testing"

func TestIsValidAuthCode(t *testing.T) {
	tests := []struct {
		name   string
		item   string
		wantOk bool
	}{
		{"微信付款码", "145007001630843683", true},
		{"异常付款码", "165007001630843683", false},
		{"支付宝付款码", "2723300600872357392234", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := IsValidAuthCode(tt.item); gotOk != tt.wantOk {
				t.Errorf("IsValidAuthCode = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
