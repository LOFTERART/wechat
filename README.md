# wechat

[![Build Status](https://travis-ci.org/cuckoopark/wechat.svg?branch=master)](https://travis-ci.org/cuckoopark/wechat)
[![Latest Tag](https://img.shields.io/github/tag/cuckoopark/wechat.svg)](https://github.com/cuckoopark/wechat/releases/latest)

这是用Golang封装了微信支付的所有API接口的SDK，并自动生成和解析XML数据。

* 支持境内普通商户和境内服务商(境外未测试)。
* 支持全局配置应用ID、商家ID等信息。
* 全部参数和返回值均使用`struct`类型传递，而不是`map`类型。

### 安装

```shell
go get -u github.com/cuckoopark/wechat
```

### 初始化

```go
const (
    isProd      = true                             // 生产环境或沙盒环境
    serviceType = wechat.ServiceTypeNormalDomestic // 普通商户或服务商等类型
    apiKey      = "xxxxxxxx"                       // 微信支付上设置的API Key
)
config := wechat.Config{
    AppId: AppID,
    SubAppId: SubAppId, // 仅服务商模式有效
    MchId: MchID,
    SubMchId: SubMchID, // 仅服务商模式有效
}
client := wechat.NewClient(isProd, serviceType, apiKey, config)
```

### 使用

下面是通用的接口，其中`client`是上面初始化时生成的实例：

* 提交付款码支付：`client.Micropay`
* 统一下单：`client.UnifiedOrder`
* 查询订单：`client.QueryOrder`
* 关闭订单：`client.CloseOrder`
* 撤销订单：client.Reverse()
* 申请退款：client.Refund()
* 查询退款：`client.QueryRefund`
* 下载对账单：`client.DownloadBill`
* 交易保障(JSAPI)：`client.ReportJsApi`
* 交易保障(MICROPAY)：`client.ReportMicropay`
* 下载资金账单：client.DownloadFundFlow()
* 拉取订单评价数据：client.BatchQueryComment()

注意事项：

* 接口调用格式为`rsp, err := client.XXX(body)`。
* 参数或返回值的类型，请查看接口对应的`wx_xxxxxx.go`文件，里面有`XXXBody`和`XXXResponse`与之对应。* 
* 参数或返回值中的常量，请参照[constant.go](constant.go)文件。
* 具体使用方法，请参照接口对应的`wx_xxxxxx_test.go`测试文件。

### 文档

* 微信支付文档：[https://pay.weixin.qq.com/wiki/doc/api/index.html](https://pay.weixin.qq.com/wiki/doc/api/index.html)
* 随机数生成算法：[https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_3](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_3)
* 签名生成算法：[https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_3](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_3)
* 交易金额：[https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2)
* 交易类型：[https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2)
* 货币类型：[https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2)
* 时间规则：[https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2)
* 时间戳：[https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2)
* 商户订单号：[https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2)
* 银行类型：[https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2)
* 单品优惠功能字段：[https://pay.weixin.qq.com/wiki/doc/api/danpin.php?chapter=9_101&index=1](https://pay.weixin.qq.com/wiki/doc/api/danpin.php?chapter=9_101&index=1)
* 代金券或立减优惠：[https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=12_1](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=12_1)
* 最新县及县以上行政区划代码：[https://pay.weixin.qq.com/wiki/doc/api/download/store_adress.csv](https://pay.weixin.qq.com/wiki/doc/api/download/store_adress.csv)


## 微信公共API

* gopay.ParseNotifyResult() => 解析并返回微信支付异步通知的参数
* gopay.VerifyPayResultSign() => 微信支付异步通知的签名验证和返回参数验签后的Sign
* gopay.Code2Session() => 登录凭证校验：获取微信用户OpenId、UnionId、SessionKey
* gopay.GetAccessToken() => 获取小程序全局唯一后台接口调用凭据
* gopay.GetPaidUnionId() => 用户支付完成后，获取该用户的 UnionId，无需用户授权
* gopay.GetWeChatUserInfo() => 微信公众号：获取用户基本信息(UnionID机制)
* gopay.DecryptOpenDataToStruct() => 加密数据，解密到指定结构体
* gopay.GetOpenIdByAuthCode() => 授权码查询openid

### 获取微信用户OpenId、UnionId、SessionKey

```go
userIdRsp, err := gopay.Code2Session(appID, secretKey, "")
if err != nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println("OpenID:", userIdRsp.Openid)
fmt.Println("UnionID:", userIdRsp.Unionid)
fmt.Println("SessionKey:", userIdRsp.SessionKey)
```

### 微信小程序支付，统一下单后，需要进一步获取微信小程序支付所需要的paySign

* 小程序支付所需要的参数，paySign由后端计算
    * timeStamp
    * nonceStr
    * package 
    * signType
    * paySign

> 官方文档说明[微信小程序支付API](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/payment/wx.requestPayment.html)
```go
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
packages := "prepay_id=" + wxRsp.PrepayId   //此处的 wxRsp.PrepayId ,统一下单成功后得到
paySign := gopay.GetMiniPaySign("wxd678efh567hg6787", wxRsp.NonceStr, packages, gopay.SignTypeMD5, timeStamp, "192006250b4c09247ec02edce69f6a2d")

//微信小程序支付需要的参数信息
fmt.Println("timeStamp：", timeStamp)
fmt.Println("nonceStr：", wxRsp.NonceStr)
fmt.Println("package：", packages)
fmt.Println("signType：", gopay.SignTypeMD5)
fmt.Println("paySign：", paySign)
```

### 微信内H5支付，统一下单后，需要进一步获取H5支付所需要的paySign

* 微信内H5支付所需要的参数，paySign由后端计算
    * appId
    * timeStamp
    * nonceStr
    * package 
    * signType
    * paySign
> 官方文档说明[微信内H5支付文档](https://pay.weixin.qq.com/wiki/doc/api/external/jsapi.php?chapter=7_7&index=6)
```go
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
packages := "prepay_id=" + wxRsp.PrepayId   //此处的 wxRsp.PrepayId ,统一下单成功后得到
paySign := gopay.GetH5PaySign("wxd678efh567hg6787", wxRsp.NonceStr, packages, gopay.SignTypeMD5, timeStamp, "192006250b4c09247ec02edce69f6a2d")

//微信内H5支付需要的参数信息
fmt.Println("appId:","wxd678efh567hg6787")
fmt.Println("timeStamp：", timeStamp)
fmt.Println("nonceStr：", wxRsp.NonceStr)
fmt.Println("package：", packages)
fmt.Println("signType：", gopay.SignTypeMD5)
fmt.Println("paySign：", paySign)
```

### APP支付，统一下单后，需要进一步获取APP支付所需要的paySign

* APP支付所需要的参数，paySign由后端计算
    * appid
    * partnerid
    * noncestr
    * prepayid
    * package 
    * timestamp
    * sign
> 官方文档说明[APP端调起支付的参数列表文档](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12)
```go
timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
//注意：signType：此处签名方式，务必与统一下单时用的签名方式一致
//注意：package：参数因为是固定值，不需开发者再传入
sign := gopay.GetH5PaySign(appid, partnerid, wxRsp.NonceStr, prepayid, gopay.SignTypeMD5, timeStamp, apiKey)

//APP支付需要的参数信息
fmt.Println("sign：", sign)
```

### 1、支付结果异步通知参数解析；2、参数解析和Sign值的验证

> 微信支付后的异步通知文档[支付结果通知](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_7&index=8)

```go
//解析支付完成后的异步通知参数信息
//此处 c.Request() 为 *http.Request
notifyRsp, err := gopay.ParseNotifyResult(c.Request())
if err != nil {
    fmt.Println("err:", err)
}
fmt.Println("notifyRsp:", notifyRsp)

//支付通知的签名验证和参数签名后的Sign
//    apiKey：API秘钥值
//    signType：签名类型 MD5 或 HMAC-SHA256（默认请填写 MD5）
//    notifyRsp：利用 gopay.ParseNotifyResult() 得到的结构体
//    返回参数ok：是否验证通过
//    返回参数sign：根据参数计算的sign值，非微信返回参数中的Sign
ok, sign := gopay.VerifyPayResultSign("192006250b4c09247ec02edce69f6a2d", "MD5", notifyRsp)
log.Println("ok:", ok)
log.Println("sign:", sign)
```

### 加密数据，解密到指定结构体

> 拿小程序获取手机号为例

button按钮获取手机号码:[button组件文档](https://developers.weixin.qq.com/miniprogram/dev/component/button.html)

微信解密算法文档:[解密算法文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)
```go
data := "Kf3TdPbzEmhWMuPKtlKxIWDkijhn402w1bxoHL4kLdcKr6jT1jNcIhvDJfjXmJcgDWLjmBiIGJ5acUuSvxLws3WgAkERmtTuiCG10CKLsJiR+AXVk7B2TUQzsq88YVilDz/YAN3647REE7glGmeBPfvUmdbfDzhL9BzvEiuRhABuCYyTMz4iaM8hFjbLB1caaeoOlykYAFMWC5pZi9P8uw=="
iv := "Cds8j3VYoGvnTp1BrjXdJg=="
sessionKey := "lyY4HPQbaOYzZdG+JcYK9w=="

phone := new(gopay.WeChatUserPhone)
err := gopay.DecryptOpenDataToStruct(data, iv, sessionKey, phone)
if err != nil {
	fmt.Println("err:", err)
	return
}
fmt.Println("PhoneNumber:", phone.PhoneNumber)
fmt.Println("PurePhoneNumber:", phone.PurePhoneNumber)
fmt.Println("CountryCode:", phone.CountryCode)
fmt.Println("Watermark:", phone.Watermark)
```

### 微信付款结果异步通知,需回复微信平台是否成功

> 代码中return写法，由于本人用的[Echo Web框架](https://github.com/labstack/echo)，有兴趣的可以尝试一下

```go
rsp := new(gopay.WeChatNotifyResponse) //回复微信的数据

rsp.ReturnCode = "SUCCESS"
rsp.ReturnMsg = "OK"

return c.String(http.StatusOK, rsp.ToXmlString())
```

### 申请退款
```go
//初始化微信客户端
//    appId：应用ID
//    mchID：商户ID
//    apiKey：API秘钥值
//    isProd：是否是正式环境
client := gopay.NewWeChatClient("wxd678efh567hg6787", "1230000109", "192006250b4c09247ec02edce69f6a2d", false)

//初始化参数结构体
body := make(gopay.BodyMap)
body.Set("out_trade_no", "MfZC2segKxh0bnJSELbvKNeH3d9oWvvQ")
body.Set("nonce_str", gopay.GetRandomString(32))
body.Set("sign_type", gopay.SignTypeMD5)
s := gopay.GetRandomString(64)
fmt.Println("s:", s)
body.Set("out_refund_no", s)
body.Set("total_fee", 101)
body.Set("refund_fee", 101)

//请求申请退款（沙箱环境下，证书路径参数可传空）
//    body：参数Body
//    certFilePath：cert证书路径
//    keyFilePath：Key证书路径
//    pkcs12FilePath：p12证书路径
wxRsp, err := client.Refund(body, "", "", "")
if err != nil {
	fmt.Println("Error:", err)
}
fmt.Println("返回值：", wxRsp)
```