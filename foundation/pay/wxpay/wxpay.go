package wxpay

import (
	"crypto/tls"
	log "github.com/Sirupsen/logrus"
)



// Client 封装微信支付接口
type Client struct {
	config *Config
	tlsConfig *tls.Config
}

// NewClientWithConfig instantiate a new Client with the specific config
func NewClientWithConfig(config *Config) *Client {
	var tlsConfig *tls.Config
	if len(config.CertificateP12) > 0 {
		// 密码为商户ID
		// 这里忽略错误, 因为没有证书只会影响高级接口, 高级接口调用时输出相应的日志即可
		var err error
		tlsConfig, err = getTLSConfig(config.CertificateP12, config.MerchantID, config.RootCAPem)
		if err != nil {
			log.Warningf("[%s:%s] Failed to init tls config, secure APIs with certificate validation would be unavailable, error: %+v",
				config.AppID, config.AppName, err)
		}
	} else {
		log.Warningf("[%s:%s] certificate was not configured, secure APIs with certificate validation would be unavailable",
			config.AppID, config.AppName)
	}

	return &Client{
		config:    config,
		tlsConfig: tlsConfig,
	}
}
