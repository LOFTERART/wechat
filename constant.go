package wechat

type BodyMap map[string]interface{}

// 微信支付的基地址
const (
	baseUrl        = "https://api.mch.weixin.qq.com/"
	baseUrlSandbox = "https://api.mch.weixin.qq.com/sandboxnew/"
)

const (
	// wxUrlUnifiedOrder        = wxBaseUrl + "pay/unifiedorder"        // 统一下单
	// wxUrlUnifiedOrderSandBox = wxBaseUrlSandbox + "pay/unifiedorder" // 统一下单(沙盒)
	// wxUrlOrderQuery          = wxBaseUrl + "pay/orderquery"          // 查询订单
	// wxUrlOrderQuerySandBox   = wxBaseUrlSandbox + "pay/orderquery"   // 查询订单(沙盒)
	//
	// wxURL_CloseOrder        = wxBaseUrl + "pay/closeorder"                  // 关闭订单
	// wxURL_Refund            = wxBaseUrl + "secapi/pay/refund"               // 申请退款
	// wxURL_Reverse           = wxBaseUrl + "secapi/pay/reverse"              // 撤销订单
	// wxURL_RefundQuery       = wxBaseUrl + "pay/refundquery"                 // 查询退款
	// wxURL_DownloadBill      = wxBaseUrl + "pay/downloadbill"                // 下载对账单
	// wxURL_DownloadFundFlow  = wxBaseUrl + "pay/downloadfundflow"            // 下载资金账单
	// wxURL_BatchQueryComment = wxBaseUrl + "billcommentsp/batchquerycomment" // 拉取订单评价数据
	//
	// // 沙盒环境
	wxURL_SanBox_GetSignKey        = baseUrlSandbox + "pay/getsignkey"
	// wxURL_SanBox_CloseOrder        = wxBaseUrlSandbox + "pay/closeorder"
	// wxURL_SanBox_Refund            = wxBaseUrlSandbox + "pay/refund"
	// wxURL_SanBox_Reverse           = wxBaseUrlSandbox + "pay/reverse"
	// wxURL_SanBox_RefundQuery       = wxBaseUrlSandbox + "pay/refundquery"
	// wxURL_SanBox_DownloadBill      = wxBaseUrlSandbox + "pay/downloadbill"
	// wxURL_SanBox_DownloadFundFlow  = wxBaseUrlSandbox + "pay/downloadfundflow"
	// wxURL_SanBox_BatchQueryComment = wxBaseUrlSandbox + "billcommentsp/batchquerycomment"

	// 支付类型
	TradeTypeApplet = "JSAPI"
	TradeTypeJsApi  = "JSAPI"
	TradeTypeApp    = "APP"
	TradeTypeH5     = "MWEB"
	TradeTypeNative = "NATIVE"
)
