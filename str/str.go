package str

import (
    "crypto/md5"
	"encoding/hex"
    "bytes"
    "unicode"
    "strconv"
    "strings"
    "fmt"
    
    "unicode/utf8"
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
例子：StrIreplace("%body%", "black", "<body text='%Body%'>")
返回："THIS"
对标PHP:Str_Replace
*/             
func StrIreplace(search, replace, subject string, count ...int) string {
	n := -1
	if len(count) > 0 {
		n = count[0]
	}
	if n == 0 {
		return subject
	}
	var (
		length      = len(search)
		searchLower = strings.ToLower(search)
	)
	for {
		originLower := strings.ToLower(subject)
		if pos := strings.Index(originLower, searchLower); pos != -1 {
			subject = subject[:pos] + replace + subject[pos+length:]
			if n--; n == 0 {
				break
			}
		} else {
			break
		}
	}
	return subject
}


/*
去除字符串首尾处的空白字符（或者其他字符）,
str   :待处理的字符串
characterMask  :可选参数，过滤字符也可由 character_mask 参数指定。一般要列出所有希望过滤的字符，也可以使用 “..” 列出一个字符范围。
例子：Trim("Hello World")
返回："HelloWorld"
此函数返回字符串 str 去除首尾空白字符后的结果。如果不指定第二个参数，trim() 将去除这些字符：
*/
func Trim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimSpace(str)
	}
	return strings.Trim(str, characterMask[0])
}

/*
删除字符串开头的空白字符（或其他字符）
str   :待处理的字符串
characterMask  :可选参数，过滤字符也可由 character_mask 参数指定。一般要列出所有希望过滤的字符，也可以使用 “..” 列出一个字符范围。
例子：Ltrim(" Hello World")
返回："HelloWorld"
该函数返回一个删除了 str 最左边的空白字符的字符串。 如果不使用第二个参数， ltrim() 仅删除以下字符
*/
func Ltrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimLeftFunc(str, unicode.IsSpace)
	}
	return strings.TrimLeft(str, characterMask[0])
}

/*
删除字符串末端的空白字符（或者其他字符）
str   :输入字符串
characterMask  :通过指定 character_mask，可以指定想要删除的字符列表。简单地列出你想要删除的全部字符。使用 .. 格式，可以指定一个范围
例子：Rtrim("Hello World ")
返回："HelloWorld"
该函数删除 str 末端的空白字符（或者其他字符）并返回。
*/
func Rtrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}
	return strings.TrimRight(str, characterMask[0])
}

/*
以千位分隔符方式格式化一个数字
number   :你要格式化的数字
decimals  :要保留的小数位数
decPoint  :指定小数点显示的字符
thousandsSep :指定千位分隔符显示的字符
例子：NumberFormat(1234.56, 2, ',', ' ')返回1 234,56
例子：NumberFormat(1234.56, 2, '.', '')返回1234.57
格式化以后的 number
*/
func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(decimals)+"F", number)
	prefix, suffix := "", ""
	if decimals > 0 {
		prefix = str[:len(str)-(decimals+1)]
		suffix = str[len(str)-decimals:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if decimals > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

/*
func StripTags(html string) string {
	var b bytes.Buffer
	s, c, i, allText := []byte(html), context{}, 0, true
	// Using the transition funcs helps us avoid mangling
	// `<div title="1>2">` or `I <3 Ponies!`.
	for i != len(s) {
		if c.delim == delimNone {
			st := c.state
			// Use RCDATA instead of parsing into JS or CSS styles.
			if c.element != elementNone && !isInTag(st) {
				st = stateRCDATA
			}
			d, nread := transitionFunc[st](c, s[i:])
			i1 := i + nread
			if c.state == stateText || c.state == stateRCDATA {
				// Emit text up to the start of the tag or comment.
				j := i1
				if d.state != c.state {
					for j1 := j - 1; j1 >= i; j1-- {
						if s[j1] == '<' {
							j = j1
							break
						}
					}
				}
				b.Write(s[i:j])
			} else {
				allText = false
			}
			c, i = d, i1
			continue
		}
		i1 := i + bytes.IndexAny(s[i:], delimEnds[c.delim])
		if i1 < i {
			break
		}
		if c.delim != delimSpaceOrTagEnd {
			// Consume any quote.
			i1++
		}
		c, i = context{state: stateTag, element: c.element}, i1
	}
	if allText {
		return html
	} else if c.state == stateText || c.state == stateRCDATA {
		b.Write(s[i:])
	}
	return b.String()
}
*/

/*
查找字符串首次出现的位置
haystack   :输入字符串
needle  :查找的字符串
offset  :如果提供了此参数，搜索会从字符串该字符数的起始位置开始统计。 如果是负数，搜索会从字符串结尾指定字符数开始。
例子：StrPos("aabc", "a")
返回：0
该函数删除 str 末端的空白字符（或者其他字符）并返回。
*/
func StrPos(haystack, needle string, offsets ...int) int {
	length := len(haystack)
    offset := 0
	if len(offsets) > 0 {
		offset = offsets[0]
	}
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

/*
查找字符串首次出现的位置（不区分大小写）
haystack   :输入字符串
needle  :查找的字符串
offset  :如果提供了此参数，搜索会从字符串该字符数的起始位置开始统计。 如果是负数，搜索会从字符串结尾指定字符数开始。
例子：StrIpos("Aabc", "a")
返回：0
该函数删除 str 末端的空白字符（或者其他字符）并返回。
*/
func StrIpos(haystack, needle string, offsets ...int) int {
	length := len(haystack)
    offset := 0
	if len(offsets) > 0 {
		offset = offsets[0]
	}
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	haystack = haystack[offset:]
	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(haystack), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

/*
字符串反转函数
str   :要反转的字符串
例子：Strrev("1234")
返回：4321
*/
func Strrev(str string) string  {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


/*
Find the first occurrence of a string
haystack   :要被截取的内容
needle   :处理的分隔符
例子：Strstr("xb@ichunt.com", "@")
返回：@ichunt.com
*/
func Strstr(haystack string, needle string) string {
	if needle == "" {
		return ""
	}
	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return ""
	}
	return haystack[idx+len([]byte(needle))-1:]
}

/*
字符转数字
char   :要转的字符
例子：Ord("\n")
返回：10
*/
func Ord(char string) int {
	r, _ := utf8.DecodeRune([]byte(char))
	return int(r)
}

