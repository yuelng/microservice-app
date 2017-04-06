package wxpay

import (
	"encoding/xml"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"strings"
)

const (
	orderRefundURI = "/secapi/pay/refund"
)

type OrderRefundRequest struct {
	BaseAPIRequest
	TransactionID string `xml:"sign_type,omitempty" key:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no" key:"out_trade_no"`
	OutRefundNo   string `xml:"out_refund_no" key:"out_refund_no"`
	TotalFee      string `xml:"total_fee" key:"total_fee"`
	RefundFee     string `xml:"refund_fee" key:"refund_fee"`
	OpUserID      string `xml:"op_user_id" key:"op_user_id"`
}

type OrderRefundResponse struct {
	BaseAPIResponse
	TransactionID     string `xml:"transaction_id" key:"transaction_id"`
	OutTradeNo        string `xml:"out_trade_no" key:"out_trade_no"`
	OutRefundNo       string `xml:"out_refund_no" key:"out_refund_no"`
	RefundID          string `xml:"refund_id" key:"refund_id"`
	RefundFee         string `xml:"refund_fee" key:"refund_fee"`
	TotalFee          string `xml:"total_fee" key:"total_fee"`
	CashFee           string `xml:"cash_fee" key:"cash_fee"`
	CashRefundFee     string `xml:"cash_refund_fee" key:"cash_refund_fee"`
	CouponRefundFee   string `xml:"coupon_refund_fee" key:"coupon_refund_fee"`
	CouponRefundCount string `xml:"coupon_refund_count" key:"coupon_refund_count"`
}

func (c *Client) OrderRefundByTradeNo(tradeNo string, amount uint32) (*OrderRefundResponse, error) {
	return c.orderRefund(&OrderRefundRequest{
		OutTradeNo:  tradeNo,
		OutRefundNo: tradeNo,
		TotalFee:    fmt.Sprint(amount),
		RefundFee:   fmt.Sprint(amount),
		OpUserID:    c.config.MerchantID(),
	})
}

func (c *Client) orderRefund(request *OrderRefundRequest) (*OrderRefundResponse, error) {
	response := &OrderRefundResponse{}
	err := c.callWXPayAPI(orderRefundURI, true, request, response)
	if err != nil {
		log.Errorln("OrderRefund callWithAPIKey error:", err)
		return nil, err
	}
	log.Infoln("OrderRefund OK:", response)
	return response, nil
}
