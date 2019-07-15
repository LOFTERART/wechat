package wechat

import "encoding/json"

// 场景信息实际对象
type SceneInfoObjModel struct {
	ID       string `json:"id"`        // 门店唯一标识
	Name     string `json:"name"`      // 门店名称
	AreaCode string `json:"area_code"` // 门店所在地行政区划码，详细见《最新县及县以上行政区划代码》，https://pay.weixin.qq.com/wiki/doc/api/download/store_adress.csv
	Address  string `json:"address"`   // 门店详细地址
}

func (m SceneInfoObjModel) IsValid() bool {
	return m.ID != ""
}

// 场景信息模型
type SceneInfoModel struct {
	SceneInfo SceneInfoObjModel `json:"-"`                    // 该字段用于上报场景信息，目前支持上报实际门店信息。
	SceneStr  string            `json:"scene_info,omitempty"` // 内部使用
}

func (m SceneInfoModel) GenerateScene() {
	if m.SceneInfo.IsValid() {
		bytes, _ := json.Marshal(m.SceneInfo)
		m.SceneStr = string(bytes)
	}
}
