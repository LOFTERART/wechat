package wechat

// 公共参数信息，自动生成不需要填写
type BodyModel struct {
	AppId    string `json:"appid"`                // 微信分配的公众账号ID(企业号corpid即为此appId)
	SubAppId string `json:"sub_appid,omitempty"`  // (服务商模式，非必填) 微信分配的子商户公众账号ID
	MchId    string `json:"mch_id"`               // 微信支付分配的商户号
	SubMchId string `json:"sub_mch_id,omitempty"` // (服务商模式) 微信支付分配的子商户号
	NonceStr string `json:"nonce_str"`            // 随机字符串，不长于32位。推荐随机数生成算法
	Sign     string `json:"sign,omitempty"`       // 签名，详见签名生成算法
	SignType string `json:"sign_type,omitempty"`  // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

// 场景信息模型
type SceneInfoModel struct {
	ID       string `json:"id"`        // 门店唯一标识
	Name     string `json:"name"`      // 门店名称
	AreaCode string `json:"area_code"` // 门店所在地行政区划码，详细见《最新县及县以上行政区划代码》
	Address  string `json:"address"`   // 门店详细地址
}

// 返回结果的通信标识
type ResponseModel struct {
	ReturnCode string `xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string `xml:"return_msg"`  // 返回信息，如非空，为错误原因：签名失败/参数格式校验错误
	RetMsg     string `xml:"retmsg"`      // 沙盒时返回的错误信息
}
