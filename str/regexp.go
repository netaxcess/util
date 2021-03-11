package str

import (
    "regexp"
)


/*
提取字符串中所有中文
str:要截取的字符串
例子:MatchCn("测试的代码2345")
*/
func MatchCn(str string) []string {
	//匹配中文的规则
    regx := regexp.MustCompile(`\p{Han}*`)
    return regx.FindAllString(str, -1)
}


/*
提取字符串所有的数字，返回数字字符串
str:要截取的字符串
例子:MatchCn("测试的代码2345 334")
返回：[]string{"2345","334"}
*/
func MatchNumber(str string) []string {
	//匹配所有数字的规则
    regex := regexp.MustCompile(`[0-9\.]+`)
    return regex.FindAllString(str, -1)
}

