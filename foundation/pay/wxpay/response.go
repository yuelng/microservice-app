package wxpay

import (
	"encoding/xml"
	"strings"
)

type APIResponseInterface interface {
	Sign() string
	IsSuccess() bool
	Error() *Error
}
// BaseAPIResponse must export for xml unmarshal
type BaseAPIResponse struct {
	XMLName xml.Name `xml:"xml"` // root element should be xml

	ReturnCode string `xml:"return_code" key:"return_code"`
	ReturnMsg  string `xml:"return_msg,omitempty" key:"return_msg"`
	AppID      string `xml:"appid" key:"appid"`
	MerchantID string `xml:"mch_id" key:"mch_id"`
	DeviceInfo string `xml:"device_info,omitempty" key:"device_info"`
	NonceStr   string `xml:"nonce_str" key:"nonce_str"`
	Sign       string `xml:"sign" key:"sign"`
	ResultCode string `xml:"result_code" key:"result_code"`
	ErrCode    string `xml:"err_code" key:"err_code"`
	ErrCodeDes string `xml:"err_code_des" key:"err_code_des"`
}

func (r *OrderRefundResponse) Sign() string {
	return r.Sign
}

func (r *OrderRefundResponse) IsSuccess() bool {
	return strings.EqualFold(r.ResultCode, "success")
}

func (r *OrderRefundResponse) Error() *Error {
	if !r.IsSuccess() {
		return GetError(r.ErrCode)
	}
	return nil
}