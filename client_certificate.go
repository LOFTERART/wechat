package wechat

import (
	"crypto/tls"
	"encoding/pem"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"golang.org/x/crypto/pkcs12"
)

func (c *Client) setCertData(certPath string) {
	if c.certData != nil && len(c.certData) > 0 {
		return
	}
	certData, err := ioutil.ReadFile(certPath)
	if err == nil {
		c.certData = certData
		c.certClient = c.buildClient()
	}
	return
}

func (c *Client) buildClient() (client *http.Client) {
	// 将pkcs12证书转成pem
	cert, err := c.pkc12ToPerm()
	if err != nil {
		return
	}
	// tls配置
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	// 带证书的客户端
	client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:     3 * time.Minute,
			TLSHandshakeTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 10 * time.Minute,
				DualStack: true,
			}).DialContext,
			TLSClientConfig:    config,
			DisableCompression: true,
		},
	}
	return
}

func (c *Client) pkc12ToPerm() (cert tls.Certificate, err error) {
	blocks, err := pkcs12.ToPEM(c.certData, c.config.MchId)
	if err != nil {
		return
	}
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	cert, err = tls.X509KeyPair(pemData, pemData)
	return
}
