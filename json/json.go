// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package json provides json operations wrapping ignoring stdlib or third-party lib json.
package json

import (
	"encoding/json"
	"io"
)

// Marshal adapts to json/encoding Marshal API.
//
// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
// Refer to https://godoc.org/encoding/json#Marshal for more information.
func JsonEncode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}



// Unmarshal adapts to json/encoding Unmarshal API
//
// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// Refer to https://godoc.org/encoding/json#Unmarshal for more information.
func JsonDecode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}



// Valid reports whether data is a valid JSON encoding.
func ValidJson(data []byte) bool {
	return json.Valid(data)
}
