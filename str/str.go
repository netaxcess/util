package str

import (
    "crypto/md5"
	"encoding/hex"
    "bytes"
    "unicode"
    "strconv"
    "strings"
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

/*
使用反斜线引用字符串
str :需要被转义的字符串
返回被转义后的内容
例子：Addslashes("Is your name O'reilly?")
*/
func Addslashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

/*
函数把包含数据的二进制字符串转换为十六进制值
str :需要转换的字符串
返回十六进制内容
例子：Bin2hex("11111001")
*/
func Bin2hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

/*
将对应ascii转换所指定的单个字符
str :需要转换的ascii码
返回ascii转换后的字符串
例子：Chr(27)，返回换行符
*/
func Chr(ascii int) string {
	return string(ascii)
}

/*
将字符串分割成小块
body :要分割的字符
chunklen :分割的尺寸
end :行尾序列符号
例子：ChunkSplit("1234", 1, "")，"1\r\n2\r\n3\r\n4\r\n"
对标PHP：chunk_split
*/
func ChunkSplit(body string, chunklen uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, erunes := []rune(body), []rune(end)
	l := uint(len(runes))
	if l <= 1 || l < chunklen {
		return body + end
	}
	ns := make([]rune, 0, len(runes)+len(erunes))
	var i uint
	for i = 0; i < l; i += chunklen {
		if i+chunklen > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunklen]...)
		}
		ns = append(ns, erunes...)
	}
	return string(ns)
}

/*
 返回字符串所用字符的信息，统计 string 中每个字节值（0..255）出现的次数，使用多种模式返回结果
str :要分割的字符
noSpace :分割的尺寸
例子：CountChars(我爱 Go Go Go, true)
返回：map[string]int{"G": 3, "o": 3, "我": 1, "爱": 1, }
对标PHP：count_chars
*/
func CountChars(str string, noSpace ...bool) map[string]int {
	m := make(map[string]int)
	countSpace := true
	if len(noSpace) > 0 && noSpace[0] {
		countSpace = false
	}
	for _, r := range []rune(str) {
		if !countSpace && unicode.IsSpace(r) {
			continue
		}
		m[string(r)]++
	}
	return m
}


/*
使用一个字符串分割另一个字符串，此函数返回由字符串组成的数组，每个元素都是 string 的一个子串，它们被字符串 delimiter 作为边界点分割出来
delimiter :边界上的分隔字符
str :要切分的的字符串
例子：Explode(" ", piece1 piece2 piece3)
返回：[]string{"piece1", "piece2", "piece3"}
*/
func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

/*
将一个一维数组的值转化为字符串
glue   :默认为空的字符串
pieces :你想要转换的数组
例子：Implode(" - ", []string{"我爱", "GoFrame"})
返回："我爱 - GoFrame"
*/
func Implode(glue string, pieces []string) string {
	var buf bytes.Buffer
	l := len(pieces)
	for _, str := range pieces {
		buf.WriteString(str)
		if l--; l > 0 {
			buf.WriteString(glue)
		}
	}
	return buf.String()
}

/*
使一个字符串的第一个字符小写,返回第一个字母小写的 str ，如果是字母的话
str   :要转换的字符串
例子：Lcfirst("This Is")
返回："this is"
*/
func Lcfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

/*
将字符串的首字母转换为大写
str   :要转换的字符串
例子：Ucfirst("this")
返回："This"
*/
func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}

/*
将字符串转化为小写
str   :要转换的字符串
例子：Strtolower("This")
返回："this"
*/
func Strtolower(str string) string {
	return strings.ToLower(str)
}

/*
将字符串转化为大写
str   :要转换的字符串
例子：Strtoupper("This")
返回："THIS"
*/
func Strtoupper(str string) string {
	return strings.ToUpper(str)
}

/*
子字符串替换,该函数返回一个字符串或者数组。该字符串或数组是将 subject 中全部的 search 都被 replace 替换之后的结果
search   :查找的目标值，也就是 needle。一个数组可以指定多个目标
replace  :search 的替换值。一个数组可以被用来指定多重替换
subject  :执行替换的数组或者字符串。也就是 haystack
count   :如果被指定，它的值将被设置为替换发生的次数
例子：StrReplace("%body%", "black", "<body text='%body%'>")
返回："THIS"
对标PHP:str_replace
*/
func StrReplace(search, replace, subject string, count ...int) string {
    n := -1
	if len(count) > 0 {
		n = count[0]
	}
	return strings.Replace(subject, search, replace, n)
}


/*
Str_Replace() 的忽略大小写版本
search   :查找的目标值，也就是 needle。一个数组可以指定多个目标
replace  :search 的替换值。一个数组可以被用来指定多重替换
subject  :执行替换的数组或者字符串。也就是 haystack
count   :如果被指定，它的值将被设置为替换发生的次数
例子：StrIreplace("%body%", "black", "<body text='%body%'>")
返回："THIS"
对标PHP:Str_Replace
*/
func StrIreplace(search, replace, subject string, count ...int) string {
    n := -1
	if len(count) > 0 {
		n = count[0]
	}
	if n == 0 {
		return origin
	}
	var (
		length      = len(search)
		searchLower = strings.ToLower(search)
	)
	for {
		originLower := strings.ToLower(origin)
		if pos := strings.Index(originLower, searchLower); pos != -1 {
			origin = origin[:pos] + replace + origin[pos+length:]
			if n--; n == 0 {
				break
			}
		} else {
			break
		}
	}
	return origin
}