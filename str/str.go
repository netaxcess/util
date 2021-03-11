package str

import (
    "crypto/md5"
	"encoding/hex"
    "bytes"
)


/*
字符串截取
s:要截取的字符串
start :开始截取的位置
length:截取的长度
返回截取后的字符串
例子:Substr("测试的代码","1","2")
*/
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

/*
返回字符串的MD5加密方式
str: 要加密的字符串
返回MD5后的字符串
例子:Md5("123456")
*/
func Md5(str string) string {
	buf := []byte(str)
	hash := md5.New()
	hash.Write(buf)
	return hex.EncodeToString(hash.Sum(nil))
}

/*
字符串连接函数
strings :多个要拼接的字符串
返回字符串拼接内容
例子：Concat("abc","def","cfg")
*/
func Concat(strings ...string) string {
	var buffer bytes.Buffer
	for _, s := range strings {
		buffer.WriteString(s)
	}
	return buffer.String()
}
