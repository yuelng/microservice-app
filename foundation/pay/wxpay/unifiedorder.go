package wxpay

import (
	log "github.com/Sirupsen/logrus"
)

const (
	unifiedOrderURI = "/pay/unifiedorder"
)

type UnifiedOrderRequest struct {
	BaseAPIRequest
	DeviceInfo     string `xml:"device_info,omitempty" key:"device_info"`
	SignType       string `xml:"sign_type,omitempty" key:"sign_type"`
	Body           string `xml:"body" key:"body"`
	Detail         string `xml:"detail,omitempty" key:"detail"`
	Attach         string `xml:"attach,omitempty" key:"attach"`
	OutTradeNo     string `xml:"out_trade_no" key:"out_trade_no"`
	FeeType        string `xml:"fee_type,omitempty" key:"fee_type"`
	TotalFee       string    `xml:"total_fee" key:"total_fee"`
	SPBillCreateIP string `xml:"spbill_create_ip" key:"spbill_create_ip"`
	TimeStart      string `xml:"time_start,omitempty" key:"time_start"`
	TimeExpire     string `xml:"time_expire,omitempty" key:"time_expire"`
	GoodsTag       string `xml:"goods_tag,omitempty" key:"goods_tag"`
	NotifyURL      string `xml:"notify_url" key:"notify_url"`
	TradeType      string `xml:"trade_type" key:"trade_type"`
	ProductID      string `xml:"product_id,omitempty" key:"product_id"`
	LimitPay       string `xml:"limit_pay,omitempty" key:"limit_pay"`
	OpenID         string `xml:"openid,omitempty" key:"openid"`
}

type unifiedOrderResponse struct {
	BaseAPIResponse

	TradeType string `xml:"trade_type" key:"trade_type"`
	PrepayID  string `xml:"prepay_id" key:"prepay_id"`
	CodeURL   string `xml:"code_url" key:"code_url"`
}

func (c *Client) UnifiedOrder(request *UnifiedOrderRequest) (string, string, string, error) {
	if len(request.TradeType) == 0 {
		request.TradeType = c.config.ClientType
	}
	if len(request.SignType) == 0 {
		request.SignType = "MD5"
	}
	request.Body = c.config.AppName + "-" + request.Body

	response := &unifiedOrderResponse{}
	err := c.callWXPayAPI(unifiedOrderURI, false, request, response)
	if err != nil {
		log.Errorln("UnifiedOrder callWithAPIKey error:", err)
		return "", "", "", err
	}
	log.Infoln("UnifiedOrder OK:", request, response)
	return response.TradeType, response.PrepayID, response.CodeURL, nil
}
