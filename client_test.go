package wechat

import (
	"fmt"
	"os"
	"testing"
)

var (
	testAppId    = os.Getenv("AppID")
	testSubAppId = os.Getenv("SubAppID")
	testMchId    = os.Getenv("MchID")
	testSubMchId = os.Getenv("SubMchID")
	testApiKey   = os.Getenv("ApiKey")
	testCertPath = os.Getenv("CertFilepath")
)

var testClient = NewPayClient(true, false, ServiceTypeFacilitatorDomestic, testApiKey, testCertPath, Config{
	AppId:    testAppId,
	SubAppId: testSubAppId,
	MchId:    testMchId,
	SubMchId: testSubMchId,
})

func TestAll(t *testing.T) {
	fmt.Println("ApiKey", testApiKey)
	fmt.Println("CertFile", testCertPath)
	fmt.Println("AppId", testAppId)
	fmt.Println("MchId", testMchId)
	if testClient.isFacilitator() {
		fmt.Println("SubAppId", testSubAppId)
		fmt.Println("SubMchId", testSubMchId)
	}
	var (
		outTradeNo    string
		transactionId string
	)
	// 付款码支付-撤销
	transactionId, outTradeNo = testMicropay(t)
	testQueryOrder(t, outTradeNo)
	testReverse(t, outTradeNo)
	// 付款码支付-关闭
	transactionId, outTradeNo = testMicropay(t)
	testQueryOrder(t, outTradeNo)
	testCloseOrder(t, outTradeNo)
	// 付款码支付-退款
	transactionId, outTradeNo = testMicropay(t)
	testQueryOrder(t, outTradeNo)
	testRefund(t, outTradeNo, transactionId)
	testQueryRefund(t, outTradeNo)
	// 统一下单支付-撤销
	outTradeNo = testUnifiedOrder(t)
	testQueryOrder(t, outTradeNo)
	testReverse(t, outTradeNo)
	// 统一下单支付-关闭
	outTradeNo = testUnifiedOrder(t)
	testQueryOrder(t, outTradeNo)
	testCloseOrder(t, outTradeNo)
	// 统一下单支付-退款
	outTradeNo = testUnifiedOrder(t)
	testQueryOrder(t, outTradeNo)
	testRefund(t, outTradeNo, "")
	testQueryRefund(t, outTradeNo)
	// 单一接口测试
	testDownloadBill(t)
	testReportMicropay(t)
}
