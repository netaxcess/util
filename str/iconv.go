package str

import (
	"github.com/axgle/mahonia"
)

/*
字符串编码转换
new_charset:要转的新编码
str        :要处理的字符
返回处理后的新编码。比如Iconv("GBK","中国人")
*/
func Iconv(new_charset,str string) string {
	return mahonia.NewDecoder(new_charset).ConvertString(str)
}