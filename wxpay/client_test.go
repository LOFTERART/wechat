package wxpay

import (
	"os"
)

var (
	testAppId    = os.Getenv("AppID")
	testSubAppId = os.Getenv("SubAppID")
	testMchId    = os.Getenv("MchID")
	testSubMchId = os.Getenv("SubMchID")
	testApiKey   = os.Getenv("ApiKey")
	testCertPath = os.Getenv("CertFilepath")
)

var testClient = NewClient(true, ServiceTypeFacilitatorDomestic, testApiKey, testCertPath, Config{
	AppId:    testAppId,
	SubAppId: testSubAppId,
	MchId:    testMchId,
	SubMchId: testSubMchId,
})
