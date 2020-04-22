/*
   Copyright 2020 XiaochengTech

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package wxpay

import (
	"crypto/tls"
	"encoding/pem"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func (c *Client) setCertData(certPath string) (err error) {
	if c.certClient != nil {
		return
	}
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return
	}
	client, err := c.buildClient(certData)
	if err != nil {
		return
	}
	c.certClient = client
	return
}

func (c *Client) buildClient(data []byte) (client *http.Client, err error) {
	// 将pkcs12证书转成pem
	cert, err := c.pkc12ToPerm(data)
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

func (c *Client) pkc12ToPerm(data []byte) (cert tls.Certificate, err error) {
	blocks, err := pkcs12.ToPEM(data, c.Config.MchId)
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
