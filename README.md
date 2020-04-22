# wechat

这是用Golang封装了微信的所有API接口的SDK。全部参数和返回值均使用`struct`类型传递，而不是`map`类型，便于使用和查错。

具体接口和使用方法，参见各目录下的说明文件：

* [微信支付](wxpay/)
* [微信服务号](wxservice/)
* [微信小程序](wxapplet/)
* [移动端APP](wxapp/)

测试方法：

修改`client_test.go`中的生成测试Client的代码，调整沙盒/生产环境、普通商户/服务商等选项，或者修改环境变量，来调整商户参数。

环境变量的脚本在`env`文件中，修改后加载环境变量：

```shell
source env
go test
```
