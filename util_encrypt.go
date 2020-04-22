package wechat

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
)

// HMAC-SHA256
func HmacSha256(str string, key string) []byte {
	hash := hmac.New(sha256.New, []byte(key))
	_, _ = hash.Write([]byte(str))
	return hash.Sum(nil)
}

// SHA1
func Sha1(str string) []byte {
	h := sha1.New()
	_, _ = h.Write([]byte(str))
	return h.Sum([]byte(""))
}

// MD5
func Md5(str string) []byte {
	hash := md5.New()
	_, _ = hash.Write([]byte(str))
	return hash.Sum(nil)
}
