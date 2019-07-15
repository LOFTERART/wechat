package wechat

import (
	"encoding/xml"
)

// 提交付款码支付
// 境内普通商户：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1
// 境内的服务商：https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_10&index=1
func (c *Client) Micropay(body MicropayBody) (wxRsp WeChatMicropayResponse, err error) {
	body.GenerateScene()
	var bytes []byte
	if bytes, err = c.doWeChat(body, c.url("pay/micropay")); err != nil {
		return
	}
	err = xml.Unmarshal(bytes, &wxRsp)
	return
}

type MicropayBody struct {
	CommonModel
	DeviceInfo     string         `json:"device_info,omitempty"` // (非必填) 终端设备号(商户自定义，如门店编号)
	Body           string         `json:"body"`                  // 商品或支付单简要描述，格式要求：门店品牌名-城市分店名-实际商品名称
	OutTradeNo     string         `json:"out_trade_no"`          // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。详见商户订单号
	TotalFee       int            `json:"total_fee"`             // 订单总金额，单位为分，只能为整数，详见支付金额
	SpbillCreateIP string         `json:"spbill_create_ip"`      // 支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP
	AuthCode       string         `json:"auth_code"`             // 扫码支付授权码，设备读取用户微信中的条码或者二维码信息 （注：用户付款码条形码规则：18位纯数字，以10、11、12、13、14、15开头）
	SceneInfoModel
}
