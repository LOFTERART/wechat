package wechat

import "os"

var (
	AppID    = os.Getenv("AppID")
	MchID    = os.Getenv("MchID")
	SubMchID = os.Getenv("SubMchID")
	ApiKey   = os.Getenv("ApiKey")
)

// 是否是生产环境
func IsProd() bool {
	return false
}

// 是普通商户还是服务商
func IsFacilitator() bool {
	return false
}
