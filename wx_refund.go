package wechat

// 申请退款 TODO
func (c *Client) Refund(body RefundBody, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp RefundResponse, err error) {
	// var bytes []byte
	// if c.isProd {
	// 	pkcsPool := x509.NewCertPool()
	// 	pkcs, iErr := ioutil.ReadFile(pkcs12FilePath)
	// 	if err = iErr; iErr != nil {
	// 		return
	// 	}
	// 	pkcsPool.AppendCertsFromPEM(pkcs)
	// 	certificate, iErr := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	// 	if err = iErr; iErr != nil {
	// 		return
	// 	}
	// 	tlsConfig := new(tls.Config)
	// 	tlsConfig.Certificates = []tls.Certificate{certificate}
	// 	tlsConfig.RootCAs = pkcsPool
	// 	tlsConfig.InsecureSkipVerify = true
	// 	if bytes, err = c.doWeChat(body, wxURL_Refund, tlsConfig); err != nil {
	// 		return
	// 	}
	// } else {
	// 	if bytes, err = c.doWeChat(body, wxURL_SanBox_Refund); err != nil {
	// 		return
	// 	}
	// }
	// err = xml.Unmarshal(bytes, &wxRsp)
	return
}

// 申请退款的参数
type RefundBody struct {
}

// 申请退款的返回值
type RefundResponse struct {
	ResponseModel
	ResultCode          string `xml:"result_code"`
	ErrCode             string `xml:"err_code"`
	ErrCodeDes          string `xml:"err_code_des"`
	Appid               string `xml:"appid"`
	MchId               string `xml:"mch_id"`
	NonceStr            string `xml:"nonce_str"`
	Sign                string `xml:"sign"`
	TransactionId       string `xml:"transaction_id"`
	OutTradeNo          string `xml:"out_trade_no"`
	OutRefundNo         string `xml:"out_refund_no"`
	RefundId            string `xml:"refund_id"`
	RefundFee           int    `xml:"refund_fee"`
	SettlementRefundFee int    `xml:"settlement_refund_fee"`
	TotalFee            int    `xml:"total_fee"`
	SettlementTotalFee  int    `xml:"settlement_total_fee"`
	FeeType             string `xml:"fee_type"`
	CashFee             int    `xml:"cash_fee"`
	CashFeeType         string `xml:"cash_fee_type"`
	CashRefundFee       int    `xml:"cash_refund_fee"`
	CouponType0         string `xml:"coupon_type_0"`
	CouponRefundFee     int    `xml:"coupon_refund_fee"`
	CouponRefundFee0    int    `xml:"coupon_refund_fee_0"`
	CouponRefundCount   int    `xml:"coupon_refund_count"`
	CouponRefundId0     string `xml:"coupon_refund_id_0"`
}
