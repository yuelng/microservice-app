package wxpay

import "encoding/xml"

type APIRequestInterface interface {
	SetAppID(appId string)
	SetMerchantID(merchantID string)
	SetNonceStr(nonceStr string)
}

type BaseAPIRequest struct {
	XMLName xml.Name `xml:"xml"` // root element should be xml

	AppID      string `xml:"appid" key:"appid"`
	MerchantID string `xml:"mch_id" key:"mch_id"`
	NonceStr   string `xml:"nonce_str" key:"nonce_str"`
	Sign       string `xml:"sign" key:"sign"`
}

func (request *BaseAPIRequest) SetAppID(appId string) {
	request.AppID = appId
}

func (request *BaseAPIRequest) SetMerchantID(merchantID string) {
	request.MerchantID = merchantID
}

func (request *BaseAPIRequest) SetNonceStr(nonceStr string) {
	request.NonceStr = nonceStr
}
