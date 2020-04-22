package wxapplet

import (
	"encoding/json"
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
)

// 获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func GetAppletUnlimitQrcode(body GetAppletUnlimitQrcodeBody) (data []byte, baseErr ResponseBase, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", body.AccessToken)
	// 参数处理
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return
	}
	params := make(map[string]interface{})
	if err = json.Unmarshal(bodyStr, &params); err != nil {
		return
	}
	if !body.AutoColor && (body.LineColorR > 0 || body.LineColorG > 0 || body.LineColorB > 0) {
		params["line_color"] = map[string]interface{}{
			"r": body.LineColorR,
			"g": body.LineColorG,
			"b": body.LineColorB,
		}
	}
	// 发送请求
	if data, err = util.HttpPostJson(url, params); err != nil {
		return
	}
	// 尝试解码
	_ = json.Unmarshal(data, &baseErr)
	return
}

type GetAppletUnlimitQrcodeBody struct {
	AccessToken string `json:"-"`                    // 接口调用凭证
	Scene       string `json:"scene"`                // 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
	Page        string `json:"page,omitempty"`       // 必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
	Width       int64  `json:"width,omitempty"`      // 二维码的宽度，单位 px，最小 280px，最大 1280px
	AutoColor   bool   `json:"auto_color,omitempty"` // 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
	LineColorR  uint8  `json:"-"`                    // auto_color为false时生效，使用rgb设置颜色
	LineColorG  uint8  `json:"-"`                    // auto_color为false时生效，使用rgb设置颜色
	LineColorB  uint8  `json:"-"`                    // auto_color为false时生效，使用rgb设置颜色
	IsHyaline   bool   `json:"is_hyaline,omitempty"` // 是否需要透明底色，为true时，生成透明底色的小程序
}
