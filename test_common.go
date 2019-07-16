package wechat

import "os"

var (
	AppID       = os.Getenv("AppID")
	MchID       = os.Getenv("MchID")
	SubMchID    = os.Getenv("SubMchID")
	ApiKey      = os.Getenv("ApiKey")
	IsProd      = false
	ServiceType = ServiceTypeNormal
)
