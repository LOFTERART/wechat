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

package util

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:     3 * time.Minute,
			TLSHandshakeTimeout: 5 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 10 * time.Minute,
				DualStack: true,
			}).DialContext,
		},
	}
}

// 发送GET请求
func HttpGet(url string) ([]byte, error) {
	rsp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)
}

// 发送POST请求(JSON)
func HttpPostJson(url string, body interface{}) ([]byte, error) {
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return httpPost(url, "application/json", string(bodyStr))
}

// 发送POST请求(XML)
func HttpPostXml(url string, xmlBody string) ([]byte, error) {
	return httpPost(url, "application/xml", xmlBody)
}

// 发送通用的POST请求
func httpPost(url string, contentType string, body string) ([]byte, error) {
	rsp, err := client.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)
}
