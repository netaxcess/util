package json

import (
	"encoding/json"
)

/*
将内容转化成json字符串
val   :需要转换的字符串
例子：JsonEncode(map[string]string{"1":"a", "2":"b"})
返回：{"1":"a","2":"b"}
*/
func JsonEncode(val interface{}) (string) {
	v, err := json.Marshal(val)
	if err != nil {
		return ""
	}

	return string(v)
}

/*
将内容转化成json字符串
val   :需要转换的字符串
v     : 回调结果map[string]interface{}
例子：JsonDecode(`{"1":"a","2":"b"}`, v)
返回：{"1":"a","2":"b"}
*/
func JsonDecode(val string) interface{} {
	var v interface{}
	json.Unmarshal([]byte(val), &v)
	return v
}



/*
验证是否正常的JSON数据
val   :需要转换的字符串
例子：ValidJson({"1":"a","2":"b"})
返回：true
*/
func ValidJson(val string) bool {
	return json.Valid([]byte(val))
}
