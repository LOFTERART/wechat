# 微信支付

* 支持境内普通商户和境内服务商(境外和银行服务商没有条件测试)。
* 支持全局配置应用ID、商家ID等信息。

### 使用方法

```go
const (
    isProd       = true                               // 生产环境或沙盒环境
    serviceType  = constant.ServiceTypeNormalDomestic // 普通商户或服务商等类型
    apiKey       = "xxxxxxxx"                         // 微信支付上设置的API Key
    certFilepath = "/xxx/yyy/apiclient_cert.p12"      // 微信证书文件的本地路径，仅部分接口使用，如果不使用这些接口，可以传递空值
)
config := wxpay.Config{
    AppId: AppID,
    MchId: MchID,
    SubAppId: SubAppId, // 仅服务商模式有效
    SubMchId: SubMchID, // 仅服务商模式有效
}
client := wxpay.NewClient(isProd, serviceType, apiKey, certFilepath, config)
```

使用初始化时生成的实例`client`进行相应支付方法的调用。例如：

```go
func Test() {
	// 初始化参数
	body := wxpay.QueryOrderBody{}
	body.OutTradeNo = "YgENQFTovdeJdFouNyy3nFVOhGD6ZvPH"
	// 请求订单查询
	wxRsp, err := client.QueryOrder(body)
	if err != nil {
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
```

注意事项：

* 参数或返回值的类型，请查看接口对应的文件，里面有`XXXBody`和`XXXResponse`与之对应。
* 参数或返回值中的常量，请参照`constant.go`文件。
* 具体使用方法，请参照接口对应的测试文件。

### 单元测试方法

修改`client_test.go`中的生成测试Client的代码，调整沙盒/生产环境、普通商户/服务商等选项，或者修改环境变量，来调整商户参数。

环境变量的脚本在`wxpay/testcase/env.profile`文件中，修改后加载环境变量：

```shell
source wxpay/testcase/env.profile
cd wxpay
go test
```

### 微信沙盒测试样例

具体说明见[wxpay/testcase](testcase/)目录。

修改`testcase/client_test.go`中的生成测试Client的代码，调整普通商户/服务商等选项，或者修改环境变量，来调整商户参数。

环境变量的脚本在`wxpay/testcase/env.profile`文件中，修改后加载环境变量：

```shell
source wxpay/testcase/env.profile
cd wxpay/testcase
go test
```

### 普通商户版API

* [ ] 付款码支付
  * [x] [付款码支付](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1)：`Micropay`
  * [x] [查询订单](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_2)：`QueryOrder`
  * [x] [撤销订单](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3)：`Reverse`
  * [x] [申请退款](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_4)：`Refund`
  * [x] [查询退款](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_5)：`QueryRefund`
  * [x] [下载交易账单](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_6)：`DownloadBill`
  * [ ] [下载资金账单](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_18&index=7)
  * [x] [交易保障](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_14&index=8)：`ReportMicropay`
  * [x] [付款码查询openid](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_13&index=9)：`AuthCodeToOpenId`
  * [x] [退款结果通知](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_16&index=10)：`NotifyRefund`
  * [ ] [拉取订单评价数据](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_17&index=11)
* [ ] JSAPI支付
  * [x] [统一下单](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1)：`UnifiedOrder`
  * [x] [查询订单](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2)：`QueryOrder`
  * [x] [关闭订单](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3)：`CloseOrder`
  * [x] [申请退款](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4)：`Refund`
  * [x] [查询退款](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5)：`QueryRefund`
  * [x] [下载交易账单](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6)：`DownloadBill`
  * [ ] [下载资金账单](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7)
  * [x] [支付结果通知](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_7&index=8)：`NotifyPay`
  * [x] [交易保障](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_8&index=9)：`ReportJsApi`
  * [x] [退款结果通知](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10)：`NotifyRefund`
  * [ ] [拉取订单评价数据](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=11)
* [ ] Native支付
* [ ] APP支付
* [ ] H5支付
* [ ] 小程序支付
  * [x] [小程序调起支付API签名](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=7_7&index=5)：`GetAppletPaySign`
  * [x] [统一下单](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1)：`UnifiedOrder`
  * [x] [查询订单](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_2)：`QueryOrder`
  * [x] [关闭订单](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_3)：`CloseOrder`
  * [x] [申请退款](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_4)：`Refund`
  * [x] [查询退款](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_5)：`QueryRefund`
  * [x] [下载交易账单](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_6)：`DownloadBill`
  * [ ] [下载资金账单](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_18&index=7)
  * [x] [支付结果通知](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_7&index=8)：`NotifyPay`
  * [x] [交易保障](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_8&index=9)：`ReportJsApi`
  * [x] [退款结果通知](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_16&index=10)：`NotifyRefund`
  * [ ] [拉取订单评价数据](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_17&index=11)
* [ ] 人脸支付
* [ ] 代金券或立减优惠
* [ ] 现金红包
* [ ] 企业付款
  * [x] 企业付款到零钱
    * [x] [企业付款](https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2)：`Change`
    * [x] [查询企业付款](https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3)：`QueryChange`
  * [ ] 企业付款到银行卡
* [ ] 分账

### 服务商版API

* [x] 付款码支付
  * [x] [付款码支付](https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_10&index=1)：`Micropay`
  * [x] [查询订单](https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_2)：`QueryOrder`
  * [x] [撤销订单](https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_11&index=3)：`Reverse`
  * [x] [申请退款](https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_4)：`Refund`
  * [x] [查询退款](https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_5)：`QueryRefund`
  * [x] [下载交易账单](https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_6)：`DownloadBill`
  * [x] [交易保障](https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_14&index=7)：`ReportMicropay`
  * [x] [付款码查询openid](https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_12&index=8)：`AuthCodeToOpenId`
  * [x] [退款结果通知](https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_16&index=9)：`NotifyRefund`
* [x] JSAPI支付
  * [x] [统一下单](https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_1)：：`UnifiedOrder`
  * [x] [查询订单](https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_2)：`QueryOrder`
  * [x] [关闭订单](https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_3)：`CloseOrder`
  * [x] [申请退款](https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_4)：`Refund`
  * [x] [查询退款](https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_5)：`QueryRefund`
  * [x] [下载交易账单](https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_6)：`DownloadBill`
  * [x] [支付结果通知](https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_7)：`NotifyPay`
  * [x] [交易保障](https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_8)：`ReportJsApi`
  * [x] [退款结果通知](https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_16)：`NotifyRefund`
* [ ] Native支付
* [ ] APP支付
* [ ] H5支付
* [x] 小程序支付
  * [x] [小程序调起支付API签名](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_sl_api.php?chapter=7_7&index=5)：`GetAppletPaySign`
  * [x] [统一下单](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_sl_api.php?chapter=9_1)：`UnifiedOrder`
  * [x] [查询订单](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_sl_api.php?chapter=9_2)：`QueryOrder`
  * [x] [关闭订单](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_sl_api.php?chapter=9_3)：`CloseOrder`
  * [x] [申请退款](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_sl_api.php?chapter=9_4)：`Refund`
  * [x] [查询退款](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_sl_api.php?chapter=9_5)：`QueryRefund`
  * [x] [下载交易账单](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_sl_api.php?chapter=9_6)：`DownloadBill`
  * [x] [支付结果通知](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_sl_api.php?chapter=9_7)：`NotifyPay`
  * [x] [交易保障](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_sl_api.php?chapter=9_8)：`ReportJsApi`
  * [x] [退款结果通知](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_sl_api.php?chapter=9_16)：`NotifyRefund`
* [ ] 人脸支付
* [ ] 现金红包
* [ ] 分账
* [ ] 特约商户进件

### 相关文档

* [普通商户版](https://pay.weixin.qq.com/wiki/doc/api/index.html)
  * 付款码支付
    * [参数规定](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_2)：包括交易金额、交易类型、货币类型、时间、时间戳、商户订单号、body字段格式、银行类型。
    * [安全规范](https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_3)：包括签名算法、生成随机数算法、API证书、商户回调API安全。
* [服务商版](https://pay.weixin.qq.com/wiki/doc/api/sl.html)
  * 付款码支付
    * [付款码支付-参数规定](https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=4_2)：包括交易金额、交易类型、货币类型、时间、时间戳、商户订单号、body字段格式、银行类型。
    * [安全规范](https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=4_3)：包括签名算法、生成随机数算法、API证书、商户回调API安全。
* 其他
  * [最新县及县以上行政区划代码](https://pay.weixin.qq.com/wiki/doc/api/download/store_adress.csv)
