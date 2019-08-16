package wechat

import (
	"crypto/tls"
	"encoding/pem"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/pkcs12"
)

func (c *Client) setCertData(certPath string) (transport *http.Transport) {
	if c.certData != nil && len(c.certData) > 0 {
		return
	}

	certData, err := ioutil.ReadFile(certPath)
	if err == nil {
		c.certData = certData
		transport = c.buildTransport()
	}
	return
}

func (c *Client) buildTransport() (transport *http.Transport) {
	// 将pkcs12证书转成pem
	cert, err := c.pkc12ToPerm()
	if err != nil {
		return
	}

	// tls配置
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	transport = &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
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
