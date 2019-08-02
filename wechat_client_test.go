/*
 普通商户测试。
*/
package wechat

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
// 	fmt.Println("返回值：", wxRsp)
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
// 	fmt.Println("返回值：", wxRsp)
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
// 	fmt.Println("返回值:", wxRsp)
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
