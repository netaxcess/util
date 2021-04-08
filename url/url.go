package url

import (
	"encoding/base64"
    "net/url"
)

/*
Encodes data with MIME base64
str   :要执行base64_encode的字符串
例子：Base64Decode("This is an encoded string")
返回："VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw=="
*/
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

/*
Decodes data encoded with MIME base64
str   :要执行base64_decode的字符串
例子：Base64Decode("VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw==")
返回："This is an encoded string"
*/
func Base64Decode(str string) (string, error) {
	switch len(str) % 4 {
	case 2:
		str += "=="
	case 3:
		str += "="
	}

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}


/*
Decodes URL-encoded string
str   :要执行url_encode的字符串
例子：URLEncode("LM358D-T%DTD")
返回："LM358D-T%25DTD"
*/
func URLEncode(str string) string {
	return url.QueryEscape(str)
}

/*
URL-encodes string
str   :要执行url_decode的字符串
例子：URLDecode("LM358D-T%25DTD")
返回："LM358D-T%25DTD"
*/
func URLDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}