package wxpay
// from wuwei

import (
	log "github.com/Sirupsen/logrus"
	"strings"
)

const (
	orderQueryURI = "/pay/orderquery"
)

type OrderQueryRequest struct {
	BaseAPIRequest
	TransactionID string `xml:"sign_type,omitempty" key:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no" key:"out_trade_no"`
}

type OrderQueryResponse struct {
	BaseAPIResponse

	OpenID             string `xml:"openid" key:"openid"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" key:"is_subscribe"`
	TradeType          string `xml:"trade_type" key:"trade_type"`
	TradeState         string `xml:"trade_state,omitempty" key:"trade_state"`
	BankType           string `xml:"bank_type,omitempty" key:"bank_type"`
	TotalFee           string `xml:"total_fee,omitempty" key:"total_fee"`
	FeeType            string `xml:"fee_type,omitempty" key:"fee_type"`
	CashFee            string `xml:"cash_fee,omitempty" key:"cash_fee"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" key:"cash_fee_type"`
	TransactionID      string `xml:"transaction_id" key:"transaction_id"`
	OutTradeNo         string `xml:"out_trade_no" key:"out_trade_no"`
	Attach             string `xml:"attach,omitempty" key:"attach"`
	TimeEnd            string `xml:"time_end" key:"time_end"`
	TradeStateDesc     string `xml:"trade_state_desc" key:"trade_state_desc"`

	//SettlementTotalFee int    `xml:"settlement_total_fee" key:"settlement_total_fee"`
	//CouponFee          int    `xml:"coupon_fee,omitempty" key:"coupon_fee"`
	//CouponCount        int    `xml:"coupon_count,omitempty" key:"coupon_count"`

	CouponBatchIDList []string
	CouponIDList      []string
	CouponFeeList     []int
}

func (r *OrderQueryResponse) IsTradeSuccess() bool {
	return strings.EqualFold(r.TradeState, "success")
}

func (c *Client) OrderQueryByTradeNo(tradeNo string) (*OrderQueryResponse, error) {
	return c.orderQueryByTradeNo(&OrderQueryRequest{
		OutTradeNo: tradeNo,
	})
}

func (c *Client) OrderQueryByTransactionID(transactionID string) (*OrderQueryResponse, error) {
	return c.orderQueryByTradeNo(&OrderQueryRequest{
		TransactionID: transactionID,
	})
}

func (c *Client) orderQueryByTradeNo(request *OrderQueryRequest) (*OrderQueryResponse, error) {
	response := &OrderQueryResponse{}
	err := c.callWXPayAPI(orderQueryURI, false, request, response)
	if err != nil {
		log.Errorln("OrderQuery callWithAPIKey error:", err)
		return nil, err
	}
	log.Infoln("OrderQuery OK:", response)
	return response, nil
}
