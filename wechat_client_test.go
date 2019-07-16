/*
 普通商户测试。
*/
package wechat

// // 测试查询订单
// func TestWeChatClientQueryOrder(t *testing.T) {
// 	client := NewWeChatClient(AppID, MchID, ApiKey, IsProd())
// 	// 初始化参数
// 	body := make(BodyMap)
// 	if IsFacilitator() {
// 		body.Set("sub_mch_id", SubMchID)
// 	}
// 	body.Set("out_trade_no", "ZhEmadCpxktbztJuf5N6PreISeLAC0y9")
// 	body.Set("nonce_str", GetRandomString(32))
// 	body.Set("sign_type", SignTypeMD5)
// 	// 请求订单查询
// 	wxRsp, err := client.QueryOrder(body)
// 	if err != nil {
// 		fmt.Printf("%s\n", err)
// 	}
// 	fmt.Printf("Response: %+v\n", wxRsp)
// }
//
// func TestWeChatClient_CloseOrder(t *testing.T) {
// 	//初始化微信客户端
// 	//    appId：应用ID
// 	//    MchID：商户ID
// 	//    ApiKey：Key值
// 	//    isProd：是否是正式环境
// 	client := NewWeChatClient(AppID, MchID, ApiKey, false)
//
// 	//初始化参数结构体
// 	body := make(BodyMap)
// 	body.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ")
// 	body.Set("nonce_str", GetRandomString(32))
// 	body.Set("sign_type", SignTypeMD5)
//
// 	//请求订单查询，成功后得到结果
// 	wxRsp, err := client.CloseOrder(body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println("Response：", wxRsp)
// }
//
// func TestWeChatClient_Refund(t *testing.T) {
// 	//初始化微信客户端
// 	//    appId：应用ID
// 	//    MchID：商户ID
// 	//    ApiKey：Key值
// 	//    isProd：是否是正式环境
// 	client := NewWeChatClient(AppID, MchID, ApiKey, true)
//
// 	//初始化参数结构体
// 	body := make(BodyMap)
// 	body.Set("out_trade_no", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk")
// 	body.Set("nonce_str", GetRandomString(32))
// 	body.Set("sign_type", SignTypeMD5)
// 	s := GetRandomString(64)
// 	fmt.Println("s:", s)
// 	body.Set("out_refund_no", s)
// 	body.Set("total_fee", 1)
// 	body.Set("refund_fee", 1)
//
// 	//请求申请退款（沙箱环境下，证书路径参数可传空）
// 	//    body：参数Body
// 	//    certFilePath：cert证书路径
// 	//    keyFilePath：Key证书路径
// 	//    pkcs12FilePath：p12证书路径
// 	wxRsp, err := client.Refund(body, Cert_iguiyu+"/apiclient_cert.pem", Cert_iguiyu+"/apiclient_key.pem", Cert_iguiyu+"/apiclient_cert.p12")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println("Response：", wxRsp)
// }
//
// func TestWeChatClient_QueryRefund(t *testing.T) {
// 	//初始化微信客户端
// 	//    appId：应用ID
// 	//    MchID：商户ID
// 	//    ApiKey：Key值
// 	//    isProd：是否是正式环境
// 	client := NewWeChatClient(AppID, MchID, ApiKey, false)
//
// 	//初始化参数结构体
// 	body := make(BodyMap)
// 	body.Set("out_refund_no", "vk4264I1UQ3Hm3E4AKsavK8npylGSgQA092f9ckUxp8A2gXmnsLEdsupURVTcaC7")
// 	body.Set("nonce_str", GetRandomString(32))
// 	body.Set("sign_type", SignTypeMD5)
//
// 	//请求申请退款
// 	wxRsp, err := client.QueryRefund(body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println("Response：", wxRsp)
// }
//
// func TestWeChatClient_DownloadBill(t *testing.T) {
// 	//初始化微信客户端
// 	//    appId：应用ID
// 	//    MchID：商户ID
// 	//    ApiKey：Key值
// 	//    isProd：是否是正式环境
// 	client := NewWeChatClient(AppID, MchID, ApiKey, false)
//
// 	//初始化参数结构体
// 	body := make(BodyMap)
// 	body.Set("nonce_str", GetRandomString(32))
// 	body.Set("sign_type", SignTypeMD5)
// 	body.Set("bill_date", "20190122")
// 	body.Set("bill_type", "ALL")
//
// 	//请求订单查询，成功后得到结果
// 	wxRsp, err := client.DownloadBill(body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println("Response：", wxRsp)
// }
//
// func TestWeChatClient_DownloadFundFlow(t *testing.T) {
// 	//初始化微信客户端
// 	//    appId：应用ID
// 	//    MchID：商户ID
// 	//    ApiKey：Key值
// 	//    isProd：是否是正式环境
// 	client := NewWeChatClient(AppID, MchID, ApiKey, false)
//
// 	//初始化参数结构体
// 	body := make(BodyMap)
// 	body.Set("nonce_str", GetRandomString(32))
// 	body.Set("sign_type", SignTypeHmacSHA256)
// 	body.Set("bill_date", "20190122")
// 	body.Set("account_type", "Basic")
//
// 	//请求订单查询，成功后得到结果，沙箱环境下，证书路径参数可传空
// 	wxRsp, err := client.DownloadFundFlow(body, "", "", "")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println("Response：", wxRsp)
// }
//
// func TestWeChatClient_BatchQueryComment(t *testing.T) {
// 	//初始化微信客户端
// 	//    appId：应用ID
// 	//    MchID：商户ID
// 	//    ApiKey：Key值
// 	//    isProd：是否是正式环境
// 	client := NewWeChatClient(AppID, MchID, ApiKey, false)
//
// 	//初始化参数结构体
// 	body := make(BodyMap)
// 	body.Set("nonce_str", GetRandomString(32))
// 	body.Set("sign_type", SignTypeHmacSHA256)
// 	body.Set("begin_time", "20190120000000")
// 	body.Set("end_time", "20190122174000")
// 	body.Set("offset", "0")
//
// 	//请求订单查询，成功后得到结果，沙箱环境下，证书路径参数可传空
// 	wxRsp, err := client.BatchQueryComment(body, "", "", "")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println("Response：", wxRsp)
// }

// func TestWeChatClient_Reverse(t *testing.T) {
// 	//初始化微信客户端
// 	//    appId：应用ID
// 	//    MchID：商户ID
// 	//    ApiKey：Key值
// 	//    isProd：是否是正式环境
// 	client := NewWeChatClient(AppID, MchID, ApiKey, false)
//
// 	//初始化参数Map
// 	body := make(BodyMap)
// 	body.Set("nonce_str", GetRandomString(32))
// 	body.Set("out_trade_no", "6aDCor1nUcAihrV5JBlI09tLvXbUp02B")
// 	body.Set("sign_type", SignTypeMD5)
//
// 	//请求撤销订单，成功后得到结果，沙箱环境下，证书路径参数可传空
// 	wxRsp, err := client.Reverse(body, "", "", "")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println("Response:", wxRsp)
// }
//
// func TestMd5(t *testing.T) {
// 	st := "appid=wxdaa2ab9ef87b5497&nonceStr=9k20rM66parD2U49&package=prepay_id=wx29164301554772fbc70d1d793335446010&signType=MD5&timeStamp=1548751382&key=GFDS8j98rewnmgl45wHTt980jg543wmg"
// 	hash := md5.New()
// 	hash.Write([]byte(st))
// 	sum := hash.Sum(nil)
// 	upper := strings.ToUpper(hex.EncodeToString(sum))
// 	fmt.Println(" ssad  ", upper)
// }
//
// func TestCode2Session(t *testing.T) {
// 	userIdRsp, err := Code2Session(AppID, APPSecret, "011EZg6p0VO47n1p2W4p0mle6p0EZg6u")
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Println("SessionKey:", userIdRsp.SessionKey)
// 	fmt.Println("ExpiresIn:", userIdRsp.ExpiresIn)
// 	fmt.Println("OpenID:", userIdRsp.Openid)
// 	fmt.Println("UnionID:", userIdRsp.Unionid)
// 	fmt.Println("Errcode:", userIdRsp.Errcode)
// 	fmt.Println("Errmsg:", userIdRsp.Errmsg)
// }
//
// func TestGetAccessToken(t *testing.T) {
// 	rsp, err := GetAccessToken(AppID, APPSecret)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Println("AccessToken:", rsp.AccessToken)
// 	fmt.Println("ExpiresIn:", rsp.ExpiresIn)
// 	fmt.Println("Errcode:", rsp.Errcode)
// 	fmt.Println("Errmsg:", rsp.Errmsg)
// }
//
// func TestGetPaidUnionId(t *testing.T) {
// 	accessToken := "21_3puo3mxoK6Ry2bR7Dh-qdn41wUP1wClwke8Zhf9a_i39jfWRq9rhNJZZZHaOt_Yad-Gp6u9_46dGW0RbIMz3nANInRI3m-1glvCnGW47v63sjYWV1iyTKOHGwDVxEv78kY-0OfkmkiIveVqAZCZaAAAQTQ"
// 	rsp, err := GetPaidUnionId(accessToken, "o0Df70MSI4Ygv2KQ2cLnoMN5CXI8", "4200000326201905256499385970")
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Println("Unionid:", rsp.Unionid)
// 	fmt.Println("Errcode:", rsp.Errcode)
// 	fmt.Println("Errmsg:", rsp.Errmsg)
// }
//
// func TestGetWeChatUserInfo(t *testing.T) {
// 	accessToken := "21_3puo3mxoK6Ry2bR7Dh-qdn41wUP1wClwke8Zhf9a_i39jfWRq9rhNJZZZHaOt_Yad-Gp6u9_46dGW0RbIMz3nANInRI3m-1glvCnGW47v63sjYWV1iyTKOHGwDVxEv78kY-0OfkmkiIveVqAZCZaAAAQTQ"
// 	userInfo, err := GetWeChatUserInfo(accessToken, OpenID)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Println("userInfo:", userInfo)
// }
//
// func TestDecryptOpenDataToStruct(t *testing.T) {
// 	data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
// 	iv := "Cds8j3VYoGvnTp1BrjXdJg=="
// 	session := "lyY4HPQbaOYzZdG+JcYK9w=="
//
// 	phone := new(WeChatUserPhone)
// 	err := DecryptOpenDataToStruct(data, iv, session, phone)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Println("PhoneNumber:", phone.PhoneNumber)
// 	fmt.Println("PurePhoneNumber:", phone.PurePhoneNumber)
// 	fmt.Println("CountryCode:", phone.CountryCode)
// 	fmt.Println("Watermark:", phone.Watermark)
// }
//
// func TestGetOpenIdByAuthCode(t *testing.T) {
// 	openIdRsp, err := GetOpenIdByAuthCode(AppID, MchID_iguiyu, "135127679952609396", ApiKey_iguiyu, GetRandomString(32))
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Println("ReturnCode:", openIdRsp.ReturnCode)
// 	fmt.Println("ReturnMsg:", openIdRsp.ReturnMsg)
// 	fmt.Println("ResultCode:", openIdRsp.ResultCode)
// 	fmt.Println("Appid:", openIdRsp.Appid)
// 	fmt.Println("MchId:", openIdRsp.MchId)
// 	fmt.Println("NonceStr:", openIdRsp.NonceStr)
// 	fmt.Println("err_code:", openIdRsp.ErrCode)
// 	fmt.Println("Openid:", openIdRsp.Openid)
// }
