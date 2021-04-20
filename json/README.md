golang处理slice或者map的公共函数,方便开发
func JsonEncode(val interface{}) (string)
func JsonDecode(val string) interface{}
func ValidJson(val string) bool
