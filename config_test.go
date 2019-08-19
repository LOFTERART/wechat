package wechat

import "os"

var testClient = NewClient(false, ServiceTypeFacilitatorDomestic, os.Getenv("ApiKey"), Config{
	AppId:    os.Getenv("AppID"),
	SubAppId: os.Getenv("SubAppID"),
	MchId:    os.Getenv("MchID"),
	SubMchId: os.Getenv("SubMchID"),
})
