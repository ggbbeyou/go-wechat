package go_wechat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	URLPayUnifiedOrder = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	URLPayQueryOrder   = "https://api.mch.weixin.qq.com/pay/orderquery"
	URLPayCloseOrder   = "https://api.mch.weixin.qq.com/pay/closeorder"

	URLQueryAccessToken    = "https://api.weixin.qq.com/cgi-bin/token"
	URLQuerySessionByCode  = "https://api.weixin.qq.com/sns/jscode2session"
	URLSendTemplateMessage = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send"
	URLGetWXACodeUnlimit   = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"

	SignTypeMD5 = "MD5"
	TradeType   = "JSAPI"
)

type Wechat struct {
	Request
	AppId       string
	AppSecret   string
	MchId       string
	NotifyUrl   string
	Key         string
	AccessToken string //如果需要用到accesstoken，需要提前调用获取token的接口，或者提前有token直接设置进来
	SignType    string
}

func NewWechat(appId, appSecret, mchId, notifyUrl, key string) *Wechat {
	return &Wechat{
		AppId:     appId,
		AppSecret: appSecret,
		MchId:     mchId,
		NotifyUrl: notifyUrl,
		Key:       key,
		SignType:  SignTypeMD5,
		Request:   NewRequest(5 * time.Second),
	}
}

/**********支付相关接口************/
// 接口文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1

//统一下单
func (w *Wechat) SendUnifiedOrder(order *UnifiedOrder) (*UnifiedOrderResp, error) {
	req := unifiedOrderReq{
		AppId:          w.AppId,
		MchId:          w.MchId,
		NotifyUrl:      w.NotifyUrl,
		TradeType:      TradeType,
		NonceStr:       RandStringBytesMaskImprSrc(16),
		SignType:       SignTypeMD5,
		OutTradeNo:     order.OutTradeNo,
		DeviceInfo:     order.DeviceInfo,
		Body:           order.Body,
		Detail:         order.Detail,
		Attach:         order.Attach,
		FeeType:        order.FeeType,
		TotalFee:       order.TotalFee,
		SpbillCreateIp: order.SpbillCreateIp,
		TimeStart:      order.TimeStart,
		TimeExpire:     order.TimeExpire,
		GoodsTag:       order.GoodsTag,
		ProductId:      order.ProductId,
		LimitPay:       order.LimitPay,
		OpenId:         order.OpenId,
	}

	sign, err := w.SignWechat(req)
	if err != nil {
		return nil, err
	}
	req.Sign = sign

	xmlBytes, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	result, err := w.Post(URLPayUnifiedOrder, "application/xml", bytes.NewReader(xmlBytes))
	if err != nil {
		return nil, err
	}

	var resp UnifiedOrderResp
	if err := decodeResponseXML(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

//查询订单
func (w *Wechat) QueryOrder(orderNo, transactionId string) (*QueryOrderResp, error) {
	req := QueryOrderReq{
		AppId:         w.AppId,
		MchId:         w.MchId,
		NonceStr:      RandStringBytesMaskImprSrc(16),
		SignType:      w.SignType,
		OutTradeNo:    orderNo,
		TransactionId: transactionId,
	}

	sign, err := w.SignWechat(req)
	if err != nil {
		return nil, err
	}
	req.Sign = sign

	xmlBytes, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	result, err := w.Post(URLPayQueryOrder, "application/xml", bytes.NewReader(xmlBytes))
	if err != nil {
		return nil, err
	}
	var resp QueryOrderResp
	if err := decodeResponseXML(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// 关闭订单
// 以下情况需要调用关单接口：商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
func (w *Wechat) CloseOrder(orderNo string) (*CloseOrderResp, error) {
	req := CloseOrderReq{
		AppId:      w.AppId,
		MchId:      w.MchId,
		NonceStr:   RandStringBytesMaskImprSrc(16),
		SignType:   SignTypeMD5,
		OutTradeNo: orderNo,
	}

	sign, err := w.SignWechat(req)
	if err != nil {
		return nil, err
	}
	req.Sign = sign

	xmlBytes, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	result, err := w.Post(URLPayCloseOrder, "application/xml", bytes.NewReader(xmlBytes))
	if err != nil {
		return nil, err
	}
	var resp CloseOrderResp
	if err := decodeResponseXML(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

/**********功能相关接口************/

// 获取小程序全局唯一后台接口调用凭据（access_token）
// 文档地址：https://developers.weixin.qq.com/miniprogram/dev/api/getAccessToken.html?search-key=getAccessToken
func (w *Wechat) QueryAccessToken() (*TokenResp, error) {
	params := map[string]string{
		"appid":      w.AppId,
		"secret":     w.AppSecret,
		"grant_type": "client_credential",
	}
	result, err := w.Get(URLQueryAccessToken, params)
	if err != nil {
		return nil, err
	}
	var resp TokenResp
	if err := decodeResponseJson(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

//小程序用`code`换`session`的接口（可以直接用小程序的云开发功能获取更方便）
func (w *Wechat) QuerySessionByCode(jsCode string) (*Session2CodeResp, error) {
	params := map[string]string{
		"appid":      w.AppId,
		"secret":     w.AppSecret,
		"js_code":    jsCode,
		"grant_type": "authorization_code",
	}
	result, err := w.Get(URLQuerySessionByCode, params)
	if err != nil {
		return nil, err
	}
	var resp Session2CodeResp
	if err := decodeResponseJson(result, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 小程序发送模板消息
// 需要注意，如果access_token过期了，返回的错误码是4001，需要自行处理，例如重发之类的
func (w *Wechat) SendTemplateMessage(req SendTemplateReq) (*SendTemplateResp, error) {
	resp := new(SendTemplateResp)
	if err := w.doPostCall(URLSendTemplateMessage, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// 获取小程序码，适用于需要的码数量极多的业务场景。
// 通过该接口生成的小程序码，永久有效，数量暂无限制
// 文档：https://developers.weixin.qq.com/miniprogram/dev/api/getWXACodeUnlimit.html?search-key=getwxacodeunlimit
func (w *Wechat) GetWXACodeUnlimit(req WXACodeUnlimitReq) (*WXACodeUnlimitResp, error) {
	resp := new(WXACodeUnlimitResp)
	if err := w.doPostCall(URLGetWXACodeUnlimit, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *Wechat) doPostCall(url string, req interface{}, resp interface{}) error {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	reqUrl := url + "?access_token=" + w.AccessToken
	result, err := w.Post(reqUrl, "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}

	if err := decodeResponseJson(result, resp); err != nil {
		return err
	}

	return nil
}

/**********工具相关方法************/
// 生成签名
func (w *Wechat) SignWechat(req interface{}) (string, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	var params map[string]string
	//由于 interface 没有办法取到 xml，导致这里Unmarshal总是会报错，导致xml这个字段没有值,所以这里先忽略...
	json.Unmarshal(reqBytes, &params)

	paramString, err := GenRequestString(params)

	if err != nil {
		return "", err
	}

	stringSignTemp := paramString + "&key=" + w.Key
	sign := HashMd5(stringSignTemp)
	return sign, nil
}

// 生成预支付数据
func (w *Wechat) GenPrepayData(prepayId, nonceStr string) (PrepayResp, error) {
	if nonceStr == "" {
		nonceStr = RandStringBytesMaskImprSrc(16)
	}
	resp := PrepayResp{
		AppId:     w.AppId,
		TimeStamp: strconv.Itoa(int(time.Now().Unix())),
		NonceStr:  nonceStr,
		Package:   "prepay_id=" + prepayId,
		SignType:  SignTypeMD5,
	}
	//计算prepay的签名
	sign, err := w.SignWechat(resp)
	if err != nil {
		return resp, err
	}
	resp.PaySign = sign

	return resp, nil
}

func (w *Wechat) ParseNotifyReq(notifyReqBytes []byte) (NotifyReq, error) {
	var req NotifyReq
	if err := xml.Unmarshal(notifyReqBytes, &req); err != nil {
		return req, err
	}

	//只有在通信成功的时候才需要校验签名,否则没有需要带校验字段
	if req.ReturnCode == "SUCCESS" {
		//验证签名
		flag := w.VerifyNotifySign(req)
		if !flag {
			return req, errors.New("verify sign failed")
		}
	}

	return req, nil
}

func (w *Wechat) GenNotifyResp(errMsg string) string {
	var code string
	if errMsg == "" {
		code = "SUCCESS"
		errMsg = "OK"
	} else {
		code = "FAIL"
	}

	resp := NotifyResp{
		ReturnCode: code,
		ReturnMsg:  errMsg,
	}
	xmlBytes, err := xml.Marshal(resp)
	if err != nil {
		return err.Error()
	}
	xmlString := string(xmlBytes)
	return xmlString
}

// 校验签名
func (w *Wechat) VerifyNotifySign(req NotifyReq) bool {
	//先取出原先的签名保存之后置空签名字段
	oldSign := req.Sign
	if oldSign == "" {
		return false
	}
	req.Sign = ""

	newSign, err := w.SignWechat(req)
	if err != nil {
		return false
	}
	return newSign == oldSign
}

// 解析http请求返回的结果
func decodeResponseJson(response *http.Response, target interface{}) error {
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		log.Printf("decode response json err: %v\n", err)
		return err
	}
	log.Printf("decodeResponseJson body: %+v", target)
	return nil
}

func decodeResponseXML(response *http.Response, target interface{}) error {
	defer response.Body.Close()
	if err := xml.NewDecoder(response.Body).Decode(target); err != nil {
		log.Printf("decode response json err: %v", err)
		return err
	}
	log.Printf("decodeResponseXML body: %+v", target)
	return nil
}
