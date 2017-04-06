package wxpay

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"encoding/json"
)

var WXPay = &_WXPay{clientMap: make(map[string]*Client)}

// Config 客户端配置
type Config struct {
	AppID      string `json:"app_id"`
	AppName    string `json:"app_name"`
	APIKey     string `json:"api_key"`
	MerchantID string `json:"merchant_id"`
	ClientType string `json:"client_type"`

	// Certificate
	CertificateP12 string `json:"certificate_p12"` // p12
	RootCAPem      string `json:"rootca_pem"`      // pem
}

type _WXPay struct {
	clientMap map[string]*Client
}
// 读取config 生成client,这里的配置是数组类型
func (wx *_WXPay) InitWithConfigFile(configFilePath string) error {
	configFile, err := os.Open(configFilePath)
	defer configFile.Close()

	if err != nil {
		log.Errorln("_WXPay --- Failed to open json file:", configFilePath, ", error:", err)
		return err
	}

	var configs []*Config
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&configs); err != nil {
		log.Errorln("_WXPay --- Failed to parse json file:", configFilePath, ", error:", err)
		return err
	}

	clients := make([]*Client, len(configs))
	for idx, cfg := range configs {
		clients[idx] = NewClientWithConfig(cfg)
	}

	wx.registerClients(clients)
	return nil
}

func (wx *_WXPay) registerClients(clients []*Client) {
	log.Infoln("WXPay: Registering clients", clients)
	for _, cli := range clients {
		wx.clientMap[cli.config.AppID] = cli
	}
}

func (wx *_WXPay) GetClient(wxAppID string) *Client {
	return wx.clientMap[wxAppID]
}
