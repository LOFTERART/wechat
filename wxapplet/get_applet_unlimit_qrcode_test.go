package wxapplet

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetAppletUnlimitQrcode(t *testing.T) {
	fmt.Println("----------获取小程序二维码----------")
	// 请求接口
	body := GetAppletUnlimitQrcodeBody{
		AccessToken: "",
		Scene:       "TJ0Rg25wM2AfFltah6XXg5PxNZoyV9D4",
	}
	data, iErr, err := GetAppletUnlimitQrcode(body)
	if err != nil {
		t.Error(err)
		return
	}
	if iErr.ErrCode > 0 {
		t.Logf("错误: %+v\n", iErr)
		return
	}
	err = ioutil.WriteFile("/Users/shallot/1.jpg", data, 0666)
	t.Logf("返回文件: ~/1.jpg\n")
}
