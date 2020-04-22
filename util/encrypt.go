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
