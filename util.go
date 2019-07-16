package wechat

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 获取随机字符串
func GetRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	b := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

func convert2String(value interface{}) (valueStr string) {
	switch v := value.(type) {
	case int:
		valueStr = Int2String(v)
	case int64:
		valueStr = Int642String(v)
	case float64:
		valueStr = Float64ToString(v)
	case float32:
		valueStr = Float32ToString(v)
	case string:
		valueStr = v
	default:
		valueStr = null
	}
	return
}

const (
	null       string = ""
	TimeLayout string = "2006-01-02 15:04:05"
)

// 解析时间
func ParseDateTime(timeStr string) (datetime time.Time) {
	datetime, _ = time.ParseInLocation(TimeLayout, timeStr, time.Local)
	return
}

// 格式化
func FormatDate(dateStr string) (formatDate string) {
	// 2020-12-30T00:00:00+08:00
	if dateStr == null {
		return null
	}
	split := strings.Split(dateStr, "T")
	formatDate = split[0]
	return
}

// Float64转字符串
//    floatNum：float64数字
//    prec：精度位数（不传则默认float数字精度）
func Float64ToString(floatNum float64, prec ...int) (floatStr string) {
	if len(prec) > 0 {
		floatStr = strconv.FormatFloat(floatNum, 'f', prec[0], 64)
		return
	}
	floatStr = strconv.FormatFloat(floatNum, 'f', -1, 64)
	return
}

// Float32转字符串
//    floatNum：float32数字
//    prec：精度位数（不传则默认float数字精度）
func Float32ToString(floatNum float32, prec ...int) (floatStr string) {
	if len(prec) > 0 {
		floatStr = strconv.FormatFloat(float64(floatNum), 'f', prec[0], 32)
		return
	}
	floatStr = strconv.FormatFloat(float64(floatNum), 'f', -1, 32)
	return
}

// Int转字符串
func Int2String(intNum int) (intStr string) {
	intStr = strconv.Itoa(intNum)
	return
}

// Int64转字符串
func Int642String(intNum int64) (int64Str string) {
	// 10, 代表10进制
	int64Str = strconv.FormatInt(intNum, 10)
	return
}

// 解密填充模式（去除补全码） PKCS7UnPadding
// 解密时，需要在最后面去掉加密时添加的填充byte
func PKCS7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unpadding := int(plainText[length-1])   // 找到Byte数组最后的填充byte
	return plainText[:(length - unpadding)] // 只截取返回有效数字内的byte数组
}
