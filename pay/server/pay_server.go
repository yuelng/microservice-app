package main

import (
	log "github.com/Sirupsen/logrus"
	"yuelng.com/microservice-go/foundation/pay/wxpay"
)

type PayServerImpl struct {
	// 这里包含config或者包含manager
	// manager 包含config?
}

func newPayServer() {
	configPath := ""
	if err := wxpay.WXPay.InitWithConfigFile(configPath); err != nil {
		log.Errorln("Failed to init WXPay:", err)
		panic(err)
	}
	// wxPay 为单例模式
	wxAppID := ""
	client := wxpay.WXPay.GetClient(wxAppID)
	client.OrderQueryByTradeNo("xxxx")
}

func (pay *PayServerImpl) WXpayUnifiedOrder() {

	// 不同的支付方式 使用不同的

	// 不需要证书
	// 参数有 appid mch_id nonce_str sign body(商品描述) out_trade_no (商户订单号) 标价金额(total_fee) 终端ip(spbill_create_ip)
	// notfiy_url (通知地址) trade_type交易类型(公众号 扫码 app支付)

	// 商品描述,商户订单号,标价金额,终端IP
	// 底层config
}

func (pay *PayServerImpl) WXpayOrderQuery() {

}

func (pay *PayServerImpl) WXpaySignature() {

}
