package array

import (
	"reflect"
)

/*
返回map的key集合
elements   :要处理的map集合
例子：ArrayKeys(map[string]string{"1":"a", "2":"b"})
返回：[]string{"1", "2"}
*/
func ArrayKeys(elements map[interface{}]interface{}) []interface{} {
	i, keys := 0, make([]interface{}, len(elements))
	for key := range elements {
		keys[i] = key
		i++
	}
	return keys
}

/*
返回map的value集合
elements   :要处理的map集合
例子：ArrayValues(map[string]string{"1":"a", "2":"b"})
返回：[]string{"a", "b"}
*/
func ArrayValues(elements map[interface{}]interface{}) []interface{} {
	i, vals := 0, make([]interface{}, len(elements))
	for _, val := range elements {
		vals[i] = val
		i++
	}
	return vals
}

/*
合并几个slice
ss   :要合并的集合
例子：ArrayMerge([]string{1,2}, []string{3,4})
返回：[]string{"1", "2", "3", "4"}
*/
func ArrayMerge(ss ...[]interface{}) []interface{} {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]interface{}, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}

/*
截取slice中的指定长度
ss   :要合并的集合
例子：ArraySlice([]string{1, 2, 3, 4, 5}, 1, 2)
返回：[]string{2, 3}
*/
func ArraySlice(s []interface{}, offset, length uint) []interface{} {
	if offset > uint(len(s)) {
		panic("offset: the offset is less than the length of s")
	}
	end := offset + length
	if end < uint(len(s)) {
		return s[offset:end]
	}
	return s[offset:]
}


/*
判断map的key是否存在
key      :要判断的值
m        :要判断的map集合
例子：ArrayKeyExists("a", map[string]string{"a":"1", "b":"2"})
返回：true
*/
func ArrayKeyExists(key interface{}, m map[interface{}]interface{}) bool {
	_, ok := m[key]
	return ok
}

/*
判断内容是否在slice中
needle   :需要判断的值
haystack :slice集合
例子：InArray(1, [2]interface{}{"a", 1})
返回：true
*/
func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type muset be slice, array or map")
	}

	return false
}