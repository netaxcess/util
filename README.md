/*
返回map的key集合
elements   :要处理的map集合
例子：ArrayKeys(map[string]string{"1":"a", "2":"b"})
返回：[]string{"1", "2"}
*/

func ArrayKeys(elements map[interface}]interface}) []interface} 




/*
返回map的value集合
elements   :要处理的map集合
例子：ArrayValues(map[string]string{"1":"a", "2":"b"})
返回：[]string{"a", "b"}
*/

func ArrayValues(elements map[interface}]interface}) []interface} 




/*
合并几个slice
ss   :要合并的集合
例子：ArrayMerge([]string{1,2}, []string{3,4})
返回：[]string{"1", "2", "3", "4"}
*/

func ArrayMerge(ss ...[]interface}) []interface} 




/*
截取slice中的指定长度
ss   :要合并的集合
例子：ArraySlice([]string{1, 2, 3, 4, 5}, 1, 2)
返回：[]string{2, 3}
*/

func ArraySlice(s []interface}, offset, length uint) []interface} 




/*
判断map的key是否存在
key      :要判断的值
m        :要判断的map集合
例子：ArrayKeyExists("a", map[string]string{"a":"1", "b":"2"})
返回：true
*/

func ArrayKeyExists(key interface}, m map[interface}]interface}) bool 




/*
判断内容是否在slice中
needle   :需要判断的值
haystack :slice集合
例子：InArray(1, [2]interface{}{"a", 1})
返回：true
*/

func InArray(needle interface}, haystack interface}) bool 




/*
新建目录
dirs：需要建立的目录文件
返回创建目录是否成功
例子：MakeDir("/data/test",0777)
*/

func MakeDir(dirs string, mode os.FileMode) bool 




/*
根据文件和目录创建
bas_dir :目录路径位置
file_name :文件路基
返回整个目录
DirNames("/home/wwwroot","2020/12")，创建/home/wwwroot/2020/12的目录
*/

func DirNames(bas_dir string,file_name string) string 




/*
递归返回子目录下面所有文件
pathname :要读取的目录地址
vals :接受返回文件目录结果的数组
返回整个目录在slice
var vals []string
例子：GetAllFile("./github.com/netaxcess/util", vals)，读取/github.com/netaxcess/util目录下面所以的文件
*/

func GetAllFile(pathname string, vals []string) ([]string, error) 




/*
将内容写入文件
filename :要写入的文件地址
data :要写入的文件内容，字符串
mode :文件权限
例子：FilePutContents("./github.com/netaxcess/util/a.txt", "内容", 0666)，读取/github.com/netaxcess/util目录下面所以的文件
*/

func FilePutContents(filename string, data string, mode os.FileMode) error 




/*
读取文件内容
filename :要读取的文件
例子：FileGetContents("./github.com/netaxcess/util/a.txt")，读取/github.com/netaxcess/util目录下面所以的文件
*/

func FileGetContents(filename string) (string, error) 




/*
将interface数据数据类型转换成字符串
data:interface{}数据类型
例子:String(1234)
*/

func String(data interface}) string 




/*
将interface转换成int
data:interface{}数据类型
例子:Int("1234")
*/

func Int(data interface}) int 




/*
将interface转换成int32
data:interface{}数据类型
例子:Int32("1234")
*/

func Int32(data interface}) int32 




/*
将interface转换成int64
data:interface{}数据类型
例子:Int64("1234")
*/

func Int64(data interface}) int64 




/*
将interface转换成float32
data:interface{}数据类型
例子:float32("1234")
*/

func Float32(data interface}) float32 




/*
将interface转换成float64
data:interface{}数据类型
例子:Float64("1234")
*/

func Float64(data interface}) float64 




/*
将内容转化成json字符串
val   :需要转换的字符串
例子：JsonEncode(map[string]string{"1":"a", "2":"b"})
返回：{"1":"a","2":"b"}
*/

func JsonEncode(val interface}) (string) 




/*
将内容转化成json字符串
val   :需要转换的字符串
v     : 回调结果map[string]interface{}
例子：JsonDecode(`{"1":"a","2":"b"}`, v)
返回：{"1":"a","2":"b"}
*/

func JsonDecode(val string) interface} 




/*
验证是否正常的JSON数据
val   :需要转换的字符串
例子：ValidJson({"1":"a","2":"b"})
返回：true
*/

func ValidJson(val string) bool 




/*
产生随机数播下随机种子
*/

func init() 




/*产生公共随机数
min最小值
max最大值
返回随机数int类型
例子:Rands(100,200)
*/

func Rands(min, max int) int 




/*
字符串编码转换
new_charset:要转的新编码
str        :要处理的字符
返回处理后的新编码。比如Iconv("GBK","中国人")
*/

func Iconv(new_charset,str string) string 




/*
提取字符串中所有中文
str:要截取的字符串
例子:MatchCn("测试的代码2345")
*/

func MatchCn(str string) []string 




/*
提取字符串所有的数字，返回数字字符串
str:要截取的字符串
例子:MatchCn("测试的代码2345 334")
返回：[]string{"2345","334"}
*/

func MatchNumber(str string) []string 




/*
字符串截取
s:要截取的字符串
start :开始截取的位置
length:截取的长度
返回截取后的字符串
例子:Substr("测试的代码","1","2")
*/

func Substr(s string, start, length int) string 




/*
返回字符串的MD5加密方式
str: 要加密的字符串
返回MD5后的字符串
例子:Md5("123456")
*/

func Md5(str string) string 




/*
字符串连接函数
strings :多个要拼接的字符串
返回字符串拼接内容
例子：Concat("abc","def","cfg")
*/

func Concat(strings ...string) string 




/*
使用反斜线引用字符串
str :需要被转义的字符串
返回被转义后的内容
例子：Addslashes("Is your name O'reilly?")
*/

func Addslashes(str string) string 




/*
函数把包含数据的二进制字符串转换为十六进制值
str :需要转换的字符串
返回十六进制内容
例子：Bin2hex("11111001")
*/

func Bin2hex(str string) (string, error) 




/*
将对应ascii转换所指定的单个字符
str :需要转换的ascii码
返回ascii转换后的字符串
例子：Chr(27)，返回换行符
*/

func Chr(ascii int) string 




/*
将字符串分割成小块
body :要分割的字符
chunklen :分割的尺寸
end :行尾序列符号
例子：ChunkSplit("1234", 1, "")，"1\r\n2\r\n3\r\n4\r\n"
对标PHP：chunk_split
*/

func ChunkSplit(body string, chunklen uint, end string) string 




/*
 返回字符串所用字符的信息，统计 string 中每个字节值（0..255）出现的次数，使用多种模式返回结果
str :要分割的字符
noSpace :分割的尺寸
例子：CountChars(我爱 Go Go Go, true)
返回：map[string]int{"G": 3, "o": 3, "我": 1, "爱": 1, }
对标PHP：count_chars
*/

func CountChars(str string, noSpace ...bool) map[string]int 




/*
使用一个字符串分割另一个字符串，此函数返回由字符串组成的数组，每个元素都是 string 的一个子串，它们被字符串 delimiter 作为边界点分割出来
delimiter :边界上的分隔字符
str :要切分的的字符串
例子：Explode(" ", piece1 piece2 piece3)
返回：[]string{"piece1", "piece2", "piece3"}
*/

func Explode(delimiter, str string) []string 




/*
将一个一维数组的值转化为字符串
glue   :默认为空的字符串
pieces :你想要转换的数组
例子：Implode(" - ", []string{"我爱", "GoFrame"})
返回："我爱 - GoFrame"
*/

func Implode(glue string, pieces []string) string 




/*
使一个字符串的第一个字符小写,返回第一个字母小写的 str ，如果是字母的话
str   :要转换的字符串
例子：Lcfirst("This Is")
返回："this is"
*/

func Lcfirst(str string) string 




/*
将字符串的首字母转换为大写
str   :要转换的字符串
例子：Ucfirst("this")
返回："This"
*/

func Ucfirst(str string) string 




/*
将字符串转化为小写
str   :要转换的字符串
例子：Strtolower("This")
返回："this"
*/

func Strtolower(str string) string 




/*
将字符串转化为大写
str   :要转换的字符串
例子：Strtoupper("This")
返回："THIS"
*/

func Strtoupper(str string) string 




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

func StrReplace(search, replace, subject string, count ...int) string 




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

func StrIreplace(search, replace, subject string, count ...int) string 




/*
去除字符串首尾处的空白字符（或者其他字符）,
str   :待处理的字符串
characterMask  :可选参数，过滤字符也可由 character_mask 参数指定。一般要列出所有希望过滤的字符，也可以使用 “..” 列出一个字符范围。
例子：Trim("Hello World")
返回："HelloWorld"
此函数返回字符串 str 去除首尾空白字符后的结果。如果不指定第二个参数，trim() 将去除这些字符：
*/

func Trim(str string, characterMask ...string) string 




/*
删除字符串开头的空白字符（或其他字符）
str   :待处理的字符串
characterMask  :可选参数，过滤字符也可由 character_mask 参数指定。一般要列出所有希望过滤的字符，也可以使用 “..” 列出一个字符范围。
例子：Ltrim(" Hello World")
返回："HelloWorld"
该函数返回一个删除了 str 最左边的空白字符的字符串。 如果不使用第二个参数， ltrim() 仅删除以下字符
*/

func Ltrim(str string, characterMask ...string) string 




/*
删除字符串末端的空白字符（或者其他字符）
str   :输入字符串
characterMask  :通过指定 character_mask，可以指定想要删除的字符列表。简单地列出你想要删除的全部字符。使用 .. 格式，可以指定一个范围
例子：Rtrim("Hello World ")
返回："HelloWorld"
该函数删除 str 末端的空白字符（或者其他字符）并返回。
*/

func Rtrim(str string, characterMask ...string) string 




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

func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string 




/*
查找字符串首次出现的位置
haystack   :输入字符串
needle  :查找的字符串
offset  :如果提供了此参数，搜索会从字符串该字符数的起始位置开始统计。 如果是负数，搜索会从字符串结尾指定字符数开始。
例子：StrPos("aabc", "a")
返回：0
该函数删除 str 末端的空白字符（或者其他字符）并返回。
*/

func StrPos(haystack, needle string, offsets ...int) int 




/*
查找字符串首次出现的位置（不区分大小写）
haystack   :输入字符串
needle  :查找的字符串
offset  :如果提供了此参数，搜索会从字符串该字符数的起始位置开始统计。 如果是负数，搜索会从字符串结尾指定字符数开始。
例子：StrIpos("Aabc", "a")
返回：0
该函数删除 str 末端的空白字符（或者其他字符）并返回。
*/

func StrIpos(haystack, needle string, offsets ...int) int 




/*
字符串反转函数
str   :要反转的字符串
例子：Strrev("1234")
返回：4321
*/

func Strrev(str string) string  




/*
Find the first occurrence of a string
haystack   :要被截取的内容
needle   :处理的分隔符
例子：Strstr("xb@ichunt.com", "@")
返回：@ichunt.com
*/

func Strstr(haystack string, needle string) string 




/*
字符转数字
char   :要转的字符
例子：Ord("\n")
返回：10
*/

func Ord(char string) int 




/*日期字符串转时间戳
date日期字符串，默认是年月日2020-02-03
返回时间戳int64
例子:StrtoTime("2020-02-02")
*/

func StrtoTime(date string) int64 




/*
获取当前系统时间戳
*/

func CurrentTime() int64 




/*
时间戳转年月日
time_unix   : 当前系统时间戳
date_formate:要转换的格式
例子：Date(1605582958,"2006/01/02 15:04:05")
例子：Date(1605582958,"2006年01月02日 15:04:05")
*/

func Date(time_unix int64, date_format string) string 




/*
获取指定某天的零点时间戳
date：时间日期
例子：DateZeroTime("2020-12-09")
*/

func DateZeroTime(date string) int64 




/*
获取指定日期是本年度第几周
date：时间日期
例子：GetWeek("2020-12-09")或者GetWeek("2020-12-09 15:04:05")
返回：year年度，week星期
*/

func GetWeek(date string) (year, week int) 




/*
Encodes data with MIME base64
str   :要执行base64_encode的字符串
例子：Base64Decode("This is an encoded string")
返回："VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw=="
*/

func Base64Encode(str string) string 




/*
Decodes data encoded with MIME base64
str   :要执行base64_decode的字符串
例子：Base64Decode("VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw==")
返回："This is an encoded string"
*/

func Base64Decode(str string) (string, error) 




/*
Decodes URL-encoded string
str   :要执行url_encode的字符串
例子：URLEncode("LM358D-T%DTD")
返回："LM358D-T%25DTD"
*/

func URLEncode(str string) string 




/*
URL-encodes string
str   :要执行url_decode的字符串
例子：URLDecode("LM358D-T%25DTD")
返回："LM358D-T%25DTD"
*/

func URLDecode(str string) (string, error) 




