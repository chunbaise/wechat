package pay

type PayReturn struct {
	XMLName struct{} `xml:"xml" structs:"-"`

	// 必选返回
	ReturnCode string `xml:"return_code" structs:"return_code"` // 返回状态码
	ReturnMsg  string `xml:"return_msg" structs:"return_msg"`   // 返回信息
}

type PayCompleteResponse struct {
	XMLName struct{} `xml:"xml" json:"-" structs:"-"`

	// 必选返回
	ReturnCode string `xml:"return_code" structs:"return_code"` // 返回状态码
	ReturnMsg  string `xml:"return_msg" structs:"return_msg"`   // 返回信息
	// 以下字段在return_code为SUCCESS的时候有返回
	Appid      string `xml:"appid" structs:"appid"`             // 小程序ID
	MchId      string `xml:"mch_id" structs:"mch_id"`           // 商户号
	DeviceInfo string `xml:"device_info" structs:"device_info"` // 设备号

	// 签名相关
	NonceStr string `xml:"nonce_str" structs:"nonce_str"` // 随机字符串
	Sign     string `xml:"sign" structs:"sign"`           // 签名
	SignType string `xml:"sign_type" structs:"sign_type"` // 签名类型

	// 业务结果相关
	ResultCode string `xml:"result_code" structs:"result_code"`   // 业务结果
	ErrCode    string `xml:"err_code" structs:"err_code"`         // 错误代码
	ErrCodeDes string `xml:"err_code_des" structs:"err_code_des"` // 错误代码描述

	Openid string `xml:"openid" structs:"openid"` // 用户标识

	IsSubscribe string `xml:"is_subscribe" structs:"is_subscribe"` // 用户是否关注公众账号
	// 交易相关
	TradeType          string `xml:"trade_type" structs:"trade_type"`                     // 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定
	BankType           string `xml:"bank_type" structs:"bank_type"`                       // 银行类型，采用字符串类型的银行标识
	TotalFee           int64  `xml:"total_fee" structs:"total_fee"`                       // 订单总金额，单位为分
	SettlementTotalFee int64  `xml:"settlement_total_fee" structs:"settlement_total_fee"` // 应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	FeeType            string `xml:"fee_type" structs:"fee_type"`                         // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFee            int64  `xml:"cash_fee" structs:"cash_fee"`                         // 现金支付金额订单现金支付金额，详见支付金额
	CashFeeType        string `xml:"cash_fee_type" structs:"cash_fee_type"`               // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型

	// 代金券相关，签名要用到，为空的，先注掉
	// CouponFee   int64  `xml:"coupon_fee" structs:"coupon_fee"`   // 代金券金额
	// CouponCount int64  `xml:"coupon_count" structs:"coupon_count"` // 代金券使用数量
	// CouponType  int64  `xml:"coupon_type" structs:"coupon_type"`  // 代金券类型
	// CouponId    string `xml:"coupon_id" structs:"coupon_id"`    // 代金券ID

	TransactionId string `xml:"transaction_id" structs:"transaction_id"` // 微信支付订单号
	OutTradeNo    string `xml:"out_trade_no" structs:"out_trade_no"`     // 商户系统的订单号，与请求一致。
	Attach        string `xml:"attach" structs:"attach"`                 // 附加数据，原样返回
	TimeEnd       string `xml:"time_end" structs:"time_end"`             // 订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
}
