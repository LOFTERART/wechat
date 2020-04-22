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
