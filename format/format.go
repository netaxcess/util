package format

import (
    "reflect"
	"strconv"
	_ "strings"
    "encoding/json"
    "fmt"
)


type apiString interface {
	String() string
}

// apiError is used for type assert api for Error().
type apiError interface {
	Error() string
}

/*
将interface数据数据类型转换成字符串
data:interface{}数据类型
例子:String(1234)
*/
func String(data interface{}) string {
	if data == nil {
		return ""
	}
	switch value := data.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	default:
		// Empty checks.
		if value == nil {
			return ""
		}
		if f, ok := value.(apiString); ok {
			// If the variable implements the String() interface,
			// then use that interface to perform the conversion
			return f.String()
		}
		if f, ok := value.(apiError); ok {
			// If the variable implements the Error() interface,
			// then use that interface to perform the conversion
			return f.Error()
		}
		// Reflect checks.
		var (
			rv   = reflect.ValueOf(value)
			kind = rv.Kind()
		)
		switch kind {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		case reflect.String:
			return rv.String()
		}
		if kind == reflect.Ptr {
			return String(rv.Elem().Interface())
		}
		// Finally we use json.Marshal to convert.
		if jsonContent, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(jsonContent)
		}
	}
}


/*
将interface转换成int
data:interface{}数据类型
例子:Int("1234")
*/
func Int(data interface{}) int {
	if data == nil {
		return 0
	}
	if v, ok := data.(int); ok {
		return v
	}
	return int(Int64(data))
}


/*
将interface转换成int32
data:interface{}数据类型
例子:Int32("1234")
*/
func Int32(data interface{}) int32 {
	if data == nil {
		return 0
	}
	if v, ok := data.(int32); ok {
		return v
	}
	return int32(Int64(data))
}

/*
将interface转换成int64
data:interface{}数据类型
例子:Int64("1234")
*/
func Int64(data interface{}) int64 {
	if data == nil {
		return 0
	}
	switch value := data.(type) {
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case uint:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		return int64(value)
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	default:
		s := String(value)
		isMinus := false
		if len(s) > 0 {
			if s[0] == '-' {
				isMinus = true
				s = s[1:]
			} else if s[0] == '+' {
				s = s[1:]
			}
		}
		// Hexadecimal
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseInt(s[2:], 16, 64); e == nil {
				if isMinus {
					return -v
				}
				return v
			}
		}
		// Octal
		if len(s) > 1 && s[0] == '0' {
			if v, e := strconv.ParseInt(s[1:], 8, 64); e == nil {
				if isMinus {
					return -v
				}
				return v
			}
		}
		// Decimal
		if v, e := strconv.ParseInt(s, 10, 64); e == nil {
			if isMinus {
				return -v
			}
			return v
		}
		// Float64
		return int64(Float64(value))
	}
}


/*
将interface转换成float32
data:interface{}数据类型
例子:float32("1234")
*/
// Float32 converts <i> to float32.
func Float32(data interface{}) float32 {
	if data == nil {
		return 0
	}
	switch value := data.(type) {
	case float32:
		return value
	case float64:
		return float32(value)
	default:
		v, _ := strconv.ParseFloat(String(data), 64)
		return float32(v)
	}
}

/*
将interface转换成float64
data:interface{}数据类型
例子:Float64("1234")
*/
// Float64 converts <i> to float64.
func Float64(data interface{}) float64 {
	if data == nil {
		return 0
	}
	switch value := data.(type) {
	case float32:
		return float64(value)
	case float64:
		return value
	default:
		v, _ := strconv.ParseFloat(String(data), 64)
		return v
	}
}
