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

package wxservice

import (
	"encoding/json"
	"fmt"
	"gitee.com/xiaochengtech/wechat/util"
)

// 获取用户基本信息(UnionId机制)
func GetBasicUserInfo(accessToken, openId, lang string) (userInfo UserInfo, err error) {
	if len(lang) <= 0 {
		lang = "zh_CN"
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=%s", accessToken, openId, lang)
	body, err := util.HttpGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &userInfo)
	return
}

type UserInfo struct {
	OpenId         string   `json:"openid"`          // 用户唯一标识
	Nickname       string   `json:"nickname"`        // 用户的昵称
	Sex            int      `json:"sex"`             // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language       string   `json:"language"`        // 用户的语言，简体中文为zh_CN
	Province       string   `json:"province"`        // 用户所在省份
	City           string   `json:"city"`            // 用户所在城市
	Country        string   `json:"country"`         // 用户所在国家
	HeadimgUrl     string   `json:"headimgurl"`      // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Privilege      []string `json:"privilege"`       // 用户特权信息
	UnionId        string   `json:"unionid"`         // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Subscribe      int      `json:"subscribe"`       // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	SubscribeTime  int      `json:"subscribe_time"`  // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	SubscribeScene string   `json:"subscribe_scene"` // 返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENEPROFILE LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_OTHERS 其他
	Remark         string   `json:"remark"`          // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupId        int      `json:"groupid"`         // 用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      []int    `json:"tagid_list"`      // 用户被打上的标签ID列表
	QrScene        int      `json:"qr_scene"`        // 二维码扫码场景（开发者自定义）
	QrSceneStr     string   `json:"qr_scene_str"`    // 二维码扫码场景描述（开发者自定义）
	ErrCode        int      `json:"errcode"`         // 错误码
	ErrMsg         string   `json:"errmsg"`          // 错误信息
}
