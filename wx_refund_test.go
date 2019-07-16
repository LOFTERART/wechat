package wechat

import (
	"testing"
)

// 测试申请退款 TODO
func TestWeChatClient_Refund(t *testing.T) {
	// // 初始化参数结构体
	// body := make(BodyMap)
	// body.Set("out_trade_no", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk")
	// body.Set("nonce_str", GetRandomString(32))
	// body.Set("sign_type", SignTypeMD5)
	// s := GetRandomString(64)
	// fmt.Println("s:", s)
	// body.Set("out_refund_no", s)
	// body.Set("total_fee", 1)
	// body.Set("refund_fee", 1)
	// // 请求申请退款（沙箱环境下，证书路径参数可传空）
	// //    body：参数Body
	// //    certFilePath：cert证书路径
	// //    keyFilePath：Key证书路径
	// //    pkcs12FilePath：p12证书路径
	// wxRsp, err := testClient.Refund(body, Cert_iguiyu+"/apiclient_cert.pem", Cert_iguiyu+"/apiclient_key.pem", Cert_iguiyu+"/apiclient_cert.p12")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println("Response：", wxRsp)
}
