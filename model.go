package go_wechat

import "encoding/xml"

// 微信支付相关
type unifiedOrderReq struct {
	XMLName        xml.Name `xml:"xml" json:"-"`
	AppId          string   `json:"appid" xml:"appid"`
	MchId          string   `json:"mch_id" xml:"mch_id"`
	OutTradeNo     string   `json:"out_trade_no" xml:"out_trade_no"`
	NonceStr       string   `json:"nonce_str" xml:"nonce_str"`
	Sign           string   `json:"sign" xml:"sign"`
	SignType       string   `json:"sign_type" xml:"sign_type"`
	DeviceInfo     string   `json:"device_info" xml:"device_info"`
	Body           string   `json:"body" xml:"body"`
	Detail         string   `json:"detail" xml:"detail"`
	Attach         string   `json:"attach" xml:"attach"`
	FeeType        string   `json:"fee_type" xml:"fee_type"`
	TotalFee       string   `json:"total_fee" xml:"total_fee"`
	SpbillCreateIp string   `json:"spbill_create_ip" xml:"spbill_create_ip"`
	TimeStart      string   `json:"time_start" xml:"time_start"`
	TimeExpire     string   `json:"time_expire" xml:"time_expire"`
	GoodsTag       string   `json:"goods_tag" xml:"goods_tag"`
	NotifyUrl      string   `json:"notify_url" xml:"notify_url"`
	TradeType      string   `json:"trade_type" xml:"trade_type"`
	ProductId      string   `json:"product_id" xml:"product_id"`
	LimitPay       string   `json:"limit_pay" xml:"limit_pay"`
	OpenId         string   `json:"openid" xml:"openid"`
}

type UnifiedOrder struct {
	OutTradeNo     string `json:"out_trade_no" xml:"out_trade_no"`
	DeviceInfo     string `json:"device_info" xml:"device_info"`
	Body           string `json:"body" xml:"body"`
	Detail         string `json:"detail" xml:"detail"`
	Attach         string `json:"attach" xml:"attach"`
	FeeType        string `json:"fee_type" xml:"fee_type"`
	TotalFee       string `json:"total_fee" xml:"total_fee"`
	SpbillCreateIp string `json:"spbill_create_ip" xml:"spbill_create_ip"`
	TimeStart      string `json:"time_start" xml:"time_start"`
	TimeExpire     string `json:"time_expire" xml:"time_expire"`
	GoodsTag       string `json:"goods_tag" xml:"goods_tag"`
	NotifyUrl      string `json:"notify_url" xml:"notify_url"`
	TradeType      string `json:"trade_type" xml:"trade_type"`
	ProductId      string `json:"product_id" xml:"product_id"`
	LimitPay       string `json:"limit_pay" xml:"limit_pay"`
	OpenId         string `json:"openid" xml:"openid"`
}

type UnifiedOrderResp struct {
	XMLName    xml.Name `xml:"xml" json:"-"`
	ReturnCode string   `xml:"return_code" json:"return_code"`
	ReturnMsg  string   `xml:"return_msg" json:"return_msg"`
	AppId      string   `xml:"appid" json:"appid"`
	MchId      string   `xml:"mch_id" json:"mch_id"`
	DeviceInfo string   `xml:"device_info" json:"device_info"`
	NonceStr   string   `xml:"nonce_str" json:"nonce_str"`
	Sign       string   `xml:"sign" json:"sign"`
	ResultCode string   `xml:"result_code" json:"result_code"`
	ErrCode    string   `xml:"err_code" json:"err_code"`
	ErrCodeDes string   `xml:"err_code_des" json:"err_code_des"`
	TradeType  string   `xml:"trade_type" json:"trade_type"`
	PrepayId   string   `xml:"prepay_id" json:"prepay_id"`
	CodeUrl    string   `xml:"code_url" json:"code_url"`
}

type queryOrderReq struct {
	XMLName       xml.Name `xml:"xml" json:"-"`
	AppId         string   `json:"appid" xml:"appid"`
	MchId         string   `json:"mch_id" xml:"mch_id"`
	OutTradeNo    string   `json:"out_trade_no" xml:"out_trade_no"`
	NonceStr      string   `json:"nonce_str" xml:"nonce_str"`
	Sign          string   `json:"sign" xml:"sign"`
	SignType      string   `json:"sign_type" xml:"sign_type"`
	TransactionId string   `json:"transaction_id" xml:"transaction_id"`
}

type QueryOrderResp struct {
	XMLName            xml.Name `xml:"xml" json:"-"`
	ReturnCode         string   `xml:"return_code" json:"return_code"`
	ReturnMsg          string   `xml:"return_msg" json:"return_msg"`
	AppID              string   `xml:"appid" json:"appid"`
	MchID              string   `xml:"mch_id" json:"mch_id"`
	NonceStr           string   `xml:"nonce_str" json:"nonce_str"`
	Sign               string   `xml:"sign" json:"sign"`
	SignType           string   `xml:"sign_type" json:"sign_type"`
	ResultCode         string   `xml:"result_code" json:"result_code"`
	ErrCode            string   `xml:"err_code" json:"err_code"`
	ErrCodeDes         string   `xml:"err_code_des" json:"err_code_des"`
	DeviceInfo         string   `xml:"device_info" json:"device_info"`
	OpenId             string   `xml:"openid" json:"openid"`
	IsSubscribe        string   `xml:"is_subscribe" json:"is_subscribe"`
	TradeType          string   `xml:"trade_type" json:"trade_type"`
	TradeState         string   `xml:"trade_state" json:"trade_state"`
	BankType           string   `xml:"bank_type" json:"bank_type"`
	TotalFee           string   `xml:"total_fee" json:"total_fee"`
	SettlementTotalFee string   `xml:"settlement_total_fee" json:"settlement_total_fee"`
	FeeType            string   `xml:"fee_type" json:"fee_type"`
	CashFee            string   `xml:"cash_fee" json:"cash_fee"`
	CashFeeType        string   `xml:"cash_fee_type" json:"cash_fee_type"`
	CouponFee          string   `xml:"coupon_fee" json:"coupon_fee"`
	CouponCount        string   `xml:"coupon_count" json:"coupon_count"`
	TransactionId      string   `xml:"transaction_id" json:"transaction_id"`
	OutTradeNo         string   `xml:"out_trade_no" json:"out_trade_no"`
	Attach             string   `xml:"attach" json:"attach"`
	TimeEnd            string   `xml:"time_end" json:"time_end"`
	TradeStateDesc     string   `xml:"trade_state_desc" json:"trade_state_desc"`
}

type closeOrderReq struct {
	XMLName    xml.Name `xml:"xml" json:"-"`
	AppId      string   `json:"appid" xml:"appid"`
	MchId      string   `json:"mch_id" xml:"mch_id"`
	OutTradeNo string   `json:"out_trade_no" xml:"out_trade_no"`
	NonceStr   string   `json:"nonce_str" xml:"nonce_str"`
	Sign       string   `json:"sign" xml:"sign"`
	SignType   string   `json:"sign_type" xml:"sign_type"`
}

type CloseOrderResp struct {
	XMLName    xml.Name `xml:"xml" json:"-"`
	ReturnCode string   `xml:"return_code" json:"return_code"`
	ReturnMsg  string   `xml:"return_msg" json:"return_msg"`
	AppId      string   `xml:"appid" json:"appid"`
	MchId      string   `xml:"mch_id" json:"mch_id"`
	DeviceInfo string   `xml:"device_info" json:"device_info"`
	NonceStr   string   `xml:"nonce_str" json:"nonce_str"`
	Sign       string   `xml:"sign" json:"sign"`
	ResultCode string   `xml:"result_code" json:"result_code"`
	ResultMsg  string   `xml:"result_msg" json:"result_msg"`
	ErrCode    string   `xml:"err_code" json:"err_code"`
	ErrCodeDes string   `xml:"err_code_des" json:"err_code_des"`
}

type PrepayResp struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

type NotifyReq struct {
	XMLName            xml.Name `xml:"xml" json:"-"`
	ReturnCode         string   `xml:"return_code" json:"return_code"`
	ReturnMsg          string   `xml:"return_msg" json:"return_msg"`
	AppID              string   `xml:"appid" json:"appid"`
	MchID              string   `xml:"mch_id" json:"mch_id"`
	DeviceInfo         string   `xml:"device_info" json:"device_info"`
	NonceStr           string   `xml:"nonce_str" json:"nonce_str"`
	Sign               string   `xml:"sign" json:"sign"`
	SignType           string   `xml:"sign_type" json:"sign_type"`
	ResultCode         string   `xml:"result_code" json:"result_code"`
	ErrCode            string   `xml:"err_code" json:"err_code"`
	ErrCodeDes         string   `xml:"err_code_des" json:"err_code_des"`
	OpenId             string   `xml:"openid" json:"openid"`
	IsSubscribe        string   `xml:"is_subscribe" json:"is_subscribe"`
	TradeType          string   `xml:"trade_type" json:"trade_type"`
	BankType           string   `xml:"bank_type" json:"bank_type"`
	TotalFee           string   `xml:"total_fee" json:"total_fee"`
	SettlementTotalFee string   `xml:"settlement_total_fee" json:"settlement_total_fee"`
	FeeType            string   `xml:"fee_type" json:"fee_type"`
	CashFee            string   `xml:"cash_fee" json:"cash_fee"`
	CashFeeType        string   `xml:"cash_fee_type" json:"cash_fee_type"`
	CouponFee          string   `xml:"coupon_fee" json:"coupon_fee"`
	CouponCount        string   `xml:"coupon_count" json:"coupon_count"`
	TransactionId      string   `xml:"transaction_id" json:"transaction_id"`
	OutTradeNo         string   `xml:"out_trade_no" json:"out_trade_no"`
	Attach             string   `xml:"attach" json:"attach"`
	TimeEnd            string   `xml:"time_end" json:"time_end"`
}

type NotifyResp struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
}

/////////////////////小程序相关//////////////////
type TokenResp struct {
	ErrCode     int64  `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type Code2SessionResp struct {
	ErrCode    int64  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}

type SendTemplateReq struct {
	ToUser          string                 `json:"touser"`
	TemplateId      string                 `json:"template_id"`
	Page            string                 `json:"page"`
	FormId          string                 `json:"form_id"`
	EmphasisKeyword string                 `json:"emphasis_keyword"`
	Data            map[string]interface{} `json:"data"`
}

type SendTemplateResp struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WXACodeUnlimitReq struct {
	Scene     string           `json:"scene"`
	Page      string           `json:"page"`
	Width     int64            `json:"width"`
	AutoColor bool             `json:"auto_color"`
	LineColor map[string]int64 `json:"line_color"`
	IsHyaline bool             `json:"is_hyaline"`
}

type WXACodeUnlimitResp struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
