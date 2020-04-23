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

package testcase

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/clbanning/mxj"
)

func CheckFields(data interface{}, fields map[string]interface{}) (err error) {
	// 将data转换到map[string]interface{}上
	dataStr, err := xml.Marshal(data)
	if err != nil {
		return
	}
	obj := make(map[string]interface{})
	if obj, err = mxj.NewMapXml(dataStr); err != nil {
		return
	}
	// 处理obj，去掉最外层
	keys := make([]string, 0)
	for key, _ := range obj {
		keys = append(keys, key)
	}
	if len(keys) != 1 {
		err = errors.New("xml转换结果异常")
		return
	}
	obj = obj[keys[0]].(map[string]interface{})
	// 遍历fields进行比对
	for key, value := range fields {
		objValue, ok := obj[key]
		if !ok {
			err = errors.New(fmt.Sprintf("%s字段不存在", key))
			return
		}
		if fmt.Sprintf("%v", value) != fmt.Sprintf("%v", objValue) {
			err = errors.New(fmt.Sprintf("%s字段的值为%v，应为%v", key, objValue, value))
			return
		}
	}
	return
}
