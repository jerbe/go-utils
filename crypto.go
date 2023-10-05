package goutils

import (
	"crypto/md5"
	"encoding/hex"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/5 19:15
  @describe :
*/

// MD5 封装的一个MD5加密字节组
func MD5(src []byte) []byte {
	data := md5.Sum(src)
	res := make([]byte, 32, 32)
	hex.Encode(res, data[:])
	return res
}

// MD5String 封装的一个MD5加密字符串的方法
func MD5String(src string) string {
	return string(MD5([]byte(src)))
}
