package wechat

import "os"

var (
	AppID       = os.Getenv("AppID")
	SubAppID    = os.Getenv("SubAppID")
	MchID       = os.Getenv("MchID")
	SubMchID    = os.Getenv("SubMchID")
	ApiKey      = os.Getenv("ApiKey")
	ServiceType = ServiceTypeNormal
)
