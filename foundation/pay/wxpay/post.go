package wxpay

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
)


const (
	apiHost       = "https://api.mch.weixin.qq.com"
	contentType   = `application/x-www-form-urlencoded;charset=UTF-8"`
)

func (c *Client) callWXPayAPI(uri string, secure bool, request APIRequestInterface, resp APIResponseInterface) error {
	request.SetAppID(c.config.AppID)
	request.SetMerchantID(c.config.MerchantID)
	request.SetNonceStr(NewNonceString())

	params, err := ToParams(request)
	if err != nil {
		return err
	}

	sign := Sign(params, c.config.APIKey)
	params = append(params, Param{"sign", sign})
	postData := []byte(params.ToXmlString())

	log.Debug(string(postData))
	data, err := c.doHttpPost(apiHost+uri, postData, true)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(data, resp)
	if err != nil {
		return err
	}

	if resp.IsSuccess() {
		ok, err := Verify(resp, c.config.APIKey, resp.Sign())
		if err != nil {
			return err
		}

		if !ok {
			return errors.New("signature error")
		}
	}
	log.Debug(resp)
	return nil
}

// doRequest post the order in xml format with a sign
func (c *Client) doHttpPost(targetUrl string, body []byte, secure bool) ([]byte, error) {
	req, err := http.NewRequest("POST", targetUrl, bytes.NewReader(body))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", contentType)

	client := http.DefaultClient
	if secure {
		if c.tlsConfig == nil {
			err := fmt.Errorf("[%s:%s] certificate was not valid, secure APIs with certificate validation is unavailable",
				c.config.AppID, c.config.AppName)
			log.Errorln(targetUrl, err)
			return []byte(""), err
		}

		client = &http.Client{
			Transport: &http.Transport{TLSClientConfig: c.tlsConfig},
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respData, nil
}
