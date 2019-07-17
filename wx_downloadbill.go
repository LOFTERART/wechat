package wechat

import (
	"encoding/xml"
	"errors"
)

// 下载对账单
func (c *Client) DownloadBill(body DownloadBillBody) (wxRsp string, failRsp *DownloadBillResponse, err error) {
	bytes, err := c.doWeChat("pay/downloadbill", body)
	if err != nil {
		return
	}
	failRsp = new(DownloadBillResponse)
	err = xml.Unmarshal(bytes, failRsp)
	if err != nil {
		return string(bytes), nil, nil
	} else {
		err = errors.New(failRsp.ReturnMsg)
		return
	}
}

// 下载对账单的参数
type DownloadBillBody struct {
	SignType string `json:"sign_type,omitempty"` // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	BillDate string `json:"bill_date"`           // 下载对账单的日期，格式：20140603
	BillType string `json:"bill_type,omitempty"` // ALL，返回当日所有订单信息，默认值 SUCCESS，返回当日成功支付的订单 REFUND，返回当日退款订单 RECHARGE_REFUND，返回当日充值退款订单
	TarType  string `json:"tar_type,omitempty"`  // 非必传参数，固定值：GZIP，返回格式为.gzip的压缩包账单。不传则默认为数据流形式。
}

// 下载对账单的返回值
type DownloadBillResponse struct {
	ResponseModel
	ErrCode string `xml:"err_code"` // 失败错误码，详见错误码列表 TODO
}
