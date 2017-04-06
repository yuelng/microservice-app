package wxpay

import (
	"encoding/xml"
	"errors"
)

// TODO(wuwei): 暂不支持代金券(没解析代金券数据, 签名会校验失败)
type Notify struct {
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

	OpenID      string `xml:"openid" key:"openid"`
	IsSubscribe string `xml:"is_subscribe,omitempty" key:"is_subscribe"`
	TradeType   string `xml:"trade_type" key:"trade_type"`
	TradeState  string `xml:"trade_state,omitempty" key:"trade_state"`
	BankType    string `xml:"bank_type,omitempty" key:"bank_type"`
	TotalFee    int    `xml:"total_fee,omitempty" key:"total_fee"`
	//SettlementTotalFee int    `xml:"settlement_total_fee" key:"settlement_total_fee"`
	FeeType     string `xml:"fee_type,omitempty" key:"fee_type"`
	CashFee     int    `xml:"cash_fee" key:"cash_fee"`
	CashFeeType string `xml:"cash_fee_type,omitempty" key:"cash_fee_type"`
	//CouponFee          int    `xml:"coupon_fee,omitempty" key:"coupon_fee"`
	//CouponCount        int    `xml:"coupon_count,omitempty" key:"coupon_count"`
	TransactionID  string `xml:"transaction_id" key:"transaction_id"`
	OutTradeNo     string `xml:"out_trade_no" key:"out_trade_no"`
	Attach         string `xml:"attach,omitempty" key:"attach"`
	TimeEnd        string `xml:"time_end" key:"time_end"`
	TradeStateDesc string `xml:"trade_state_desc" key:"trade_state_desc"`

	CouponBatchIDList []string
	CouponIDList      []string
	CouponFeeList     []int
}

func NewNotifyFromBytes(bytes []byte) (*Notify, error) {
	notify := &Notify{}
	err := xml.Unmarshal(bytes, notify)
	return notify, err
}

func (c *Client) ValidateNotify(notify *Notify) error {
	ok, err := Verify(notify, c.config.APIKey, notify.Sign())
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("signature error")
	}

	return nil
}
